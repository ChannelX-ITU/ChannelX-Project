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
	ErrUserNotInGroup			= ChannelError{"Err_User_Not_In_Group", "You can not join this group with this e-mail address.", http.StatusBadRequest}
	ErrCommTypeDoesNotMatch		= ChannelError{"Err_Comm_Type_Does_Not_Match", "You can not join this channel with this communication type.", http.StatusBadRequest}
	ErrInvalidEmail			   	= ChannelError{"Err_Invalid_Email", "E-mail address is not valid", http.StatusBadRequest}
	ErrInvalidUsernameOnSignup	= ChannelError{"Err_Invalid_Username_On_Signup", "Username is not valid", http.StatusBadRequest}
	ErrInvalidPasswordOnSignup	= ChannelError{"Err_Invalid_Password_On_Signup", "Password is not valid", http.StatusBadRequest}
	ErrInvalidChannelName		= ChannelError{"Err_Invalid_Channel_Name", "Invalid channel name (Same character restrictions with username)", http.StatusBadRequest}
	ErrInvalidCommType			= ChannelError{"Err_Invalid_Comm_Type", "Invalid Communication Type", http.StatusBadRequest}
	ErrEmptyChannelPreference	= ChannelError{"Err_Empty_Channel_Preference", "Invalid Alias ", http.StatusBadRequest}
	ErrInvalidPhoneNumber		= ChannelError{"Err_Invalid_Phone_Number", "Phone number must be valid.", http.StatusBadRequest}
	ErrNotInInterval = ChannelError{"Err_Not_In_Interval", "Message can not be send because of the restriction on time interval", http.StatusPreconditionFailed}
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
	ErrCommIsTaken = ChannelError{"Err_Comm_Is_Taken", "This communication method is already taken", http.StatusBadRequest}
	ErrNoCommOfUser = ChannelError{"Err_Not_Comm_Of_User", "This communication method is not registered for the user", http.StatusBadRequest}
	ErrCommInUse = ChannelError{"Err_Comm_In_Use", "This communication method is assigned to a channel", http.StatusBadRequest}

	ErrGelbori = ChannelError{"at", "avrat", http.StatusInternalServerError}
	ErrGelbori1 = ChannelError{"at", "avrat1", http.StatusInternalServerError}
	ErrGelbori2 = ChannelError{"at", "avrat2", http.StatusInternalServerError}
	ErrGelbori3 = ChannelError{"at", "avrat3", http.StatusInternalServerError}
	ErrGelbori4 = ChannelError{"at", "avrat4", http.StatusInternalServerError}
	ErrGelbori5 = ChannelError{"at", "avrat5", http.StatusInternalServerError}
	ErrGelbori6 = ChannelError{"at", "avrat6", http.StatusInternalServerError}
	ErrGelbori7 = ChannelError{"at", "avrat7", http.StatusInternalServerError}
	ErrGelbori8 = ChannelError{"at", "avrat8", http.StatusInternalServerError}
	ErrGelbori9 = ChannelError{"at", "avrat9", http.StatusInternalServerError}


)

func WriteSuccess(w http.ResponseWriter, message string) {
	s, _ := json.Marshal(ChannelResponse{Status:"Success", Message:message, Error:nil})
	w.Write(s)
}

func WriteError(w http.ResponseWriter, err ChannelError) {
	s, _ := json.Marshal(ChannelResponse{Status:"Error", Error:&err})
	http.Error(w, string(s), err.statusCode)
}