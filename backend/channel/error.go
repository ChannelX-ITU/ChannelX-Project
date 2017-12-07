package channel

import (
	"net/http"
	"encoding/json"
)

type ChannelResponse struct {
	Status		string			`json:"status"`
	Message		string			`json:"message,omitempty"`
	Error		*ChannelError	`json:"error,omitempty"`
}

type ChannelError struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	statusCode	int
}

var (
	ErrInvalidLoginCredentials = ChannelError{"Err_Invalid_Login_Credentials", "Username or password is incorrect", http.StatusBadRequest}
	ErrUsernameTaken = ChannelError{"Err_Username_Taken", "This username is already taken", http.StatusBadRequest}
	ErrEmailTaken = ChannelError{"Err_Email_Taken", "This email is already taken", http.StatusBadRequest}
	ErrChannelNotExist = ChannelError{"Err_Channel_Not_Exist", "This channel does not exist", http.StatusBadRequest}
	ErrUserNotInChannel = ChannelError{"Err_User_Not_In_Channel", "User is not a participant in that channel", http.StatusBadRequest}
	ErrUserIsNotOwner = ChannelError{"Err_User_Is_Not_Owner", "User is not the owner of the channel", http.StatusUnauthorized}
	ErrInternalServerError = ChannelError{"Err_Internal_Server_Error", "Something went really wrong ", http.StatusInternalServerError}
	ErrWrongMethod = ChannelError{"Err_Wrong_Method", "This endpoint excepts a different http method", http.StatusMethodNotAllowed}
	ErrAccountNotActivated = ChannelError{"Err_Account_Not_Activated", "This account is not activated", http.StatusPreconditionFailed}
	ErrNotLoggedIn = ChannelError{"Err_Not_Logged_In", "This action requires login", http.StatusForbidden}
	ErrChannelExist = ChannelError{"Err_Channel_Exists", "A channel already exists with that name", http.StatusPreconditionFailed}
	ErrUserInChannel = ChannelError{"Err_User_In_Channel", "User is already a participant in that channel", http.StatusBadRequest}
	ErrCommIsTaken = ChannelError{"Err_Comm_Is_Taken", "This communication method is already registered", http.StatusBadRequest}
	ErrNoCommOfUser = ChannelError{"Err_Not_Comm_Of_User", "This communication method is not registered for the user", http.StatusBadRequest}
)

func WriteSuccess(w http.ResponseWriter, message string) {
	s, _ := json.Marshal(ChannelResponse{Status:"Success", Message:message, Error:nil})
	w.Write(s)
}

func WriteError(w http.ResponseWriter, err ChannelError) {
	s, _ := json.Marshal(ChannelResponse{Status:"Error", Error:&err})
	http.Error(w, string(s), err.statusCode)
}