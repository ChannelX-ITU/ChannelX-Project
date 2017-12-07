package channel

import (
	"regexp"
	"net/http"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var IsUserName = regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`).MatchString

var(
	ErrInvalidEmailOnSignup   	= ChannelError{"Err_Invalid_Email_On_Signup", "E-mail address is not valid", http.StatusBadRequest}
	ErrInvalidUsernameOnSignup	= ChannelError{"Err_Invalid_Username_On_Signup", "Username is not valid", http.StatusBadRequest}
	ErrInvalidPasswordOnSignup	= ChannelError{"Err_Invalid_Password_On_Signup", "Password is not valid", http.StatusBadRequest}
	ErrInvalidChannelNameLength	= ChannelError{"Err_Invalid_Channel_Name_Length", "Channel name length is not suitable", http.StatusBadRequest}
	ErrInvalidChannelName		= ChannelError{"Err_Invalid_Channel_Name", "Invalid channel name (Same character restrictions with username)", http.StatusBadRequest}
	ErrInvalidCommType			= ChannelError{"Err_Invalid_Comm_Type", "Invalid Communicaton Type", http.StatusBadRequest}
	ErrInvalidAlias				= ChannelError{"Err_Invalid_Alias", "Invalid Alias (Same character restrictions with username)", http.StatusBadRequest}
)


func ( s *Server ) IsLoginValid ( login Login ) *ChannelError {
	if IsUserName( login.Username ) == false || len(login.Password) < 8 || len(login.Password) > 16  {
		return &ErrInvalidLoginCredentials
	}
	return nil
}

func ( s *Server ) IsSignupValid ( signup SignUp ) *ChannelError {
	if validation.Validate(signup.Email, validation.Required, is.Email) != nil {
		return &ErrInvalidEmailOnSignup
	}
	if IsUserName( signup.Username ) == false {
		return &ErrInvalidUsernameOnSignup
	}
	if validation.Validate( signup.Password, validation.Required, validation.Length(6, 50) ) != nil {
		return &ErrInvalidPasswordOnSignup
	}
	return nil
}

func ( s *Server ) CanJoinChannel ( joinchannel JoinChannel ) *ChannelError {
	if len( joinchannel.Channel ) < 4 || len(joinchannel.Channel) > 30 {
		return &ErrInvalidChannelNameLength
	}
	if IsUserName( joinchannel.Channel ) {
		return &ErrInvalidChannelName
	}
	if joinchannel.Comm != `SMS`  && joinchannel.Comm != `EMAIL` {
		return &ErrInvalidCommType
	}
	if  len( joinchannel.Alias ) != 0 {
		if IsUserName(joinchannel.Alias) == false {
			return &ErrInvalidAlias
		}
	}
	return nil
}