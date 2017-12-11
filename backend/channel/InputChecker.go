package channel

import (
	"regexp"
	"net/http"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var IsUserName 		= regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)
var IsPhoneNumber	= regexp.MustCompile( `(0|\\+90|090|90)[0-9]+$`)

var(
	ErrCommTypeDoesNotMatch		= ChannelError{"Err_Comm_Type_Does_Not_Match", "You can not join this channel with this communication type.", http.StatusBadRequest}
	ErrChannelDoesNotExist		= ChannelError{"Err_Channel_Does_Not_Exist", "Channel is already deleted.", http.StatusBadRequest}
	ErrInvalidEmail			   	= ChannelError{"Err_Invalid_Email", "E-mail address is not valid", http.StatusBadRequest}
	ErrInvalidUsernameOnSignup	= ChannelError{"Err_Invalid_Username_On_Signup", "Username is not valid", http.StatusBadRequest}
	ErrInvalidPasswordOnSignup	= ChannelError{"Err_Invalid_Password_On_Signup", "Password is not valid", http.StatusBadRequest}
	ErrInvalidChannelName		= ChannelError{"Err_Invalid_Channel_Name", "Invalid channel name (Same character restrictions with username)", http.StatusBadRequest}
	ErrInvalidCommType			= ChannelError{"Err_Invalid_Comm_Type", "Invalid Communication Type", http.StatusBadRequest}
	ErrInvalidAlias				= ChannelError{"Err_Invalid_Alias", "Invalid Alias (Same character restrictions with username)", http.StatusBadRequest}
	ErrEmptyChannelPreference	= ChannelError{"Err_Empty_Channel_Preference", "Invalid Alias ", http.StatusBadRequest}
	ErrEmptyChannelRestrictions = ChannelError{"Err_Empty_Channel_Restrictions", "Channel Restrictions is not set.", http.StatusBadRequest}
	ErrInvalidPhoneNumber		= ChannelError{"Err_Invalid_Phone_Number", "Phone number must be valid.", http.StatusBadRequest}
)


func ( s *Server ) LoginValidation ( login Login ) *ChannelError {
	if validation.Validate( login.Username, validation.Required,
		validation.Match( IsUserName ), validation.Length(3, 16) ) != nil {
		return &ErrInvalidLoginCredentials
	}
	return nil
}

func ( s *Server ) SignupValidation ( signup SignUp ) *ChannelError {
	if validation.Validate(signup.Email, validation.Required, is.Email) != nil {
		return &ErrInvalidEmail
	}
	if validation.Validate( signup.Username, validation.Required,
		validation.Match( IsUserName ), validation.Length(3, 16) ) != nil {
		return &ErrInvalidUsernameOnSignup
	}
	if validation.Validate( signup.Password, validation.Required, validation.Length(6, 50) ) != nil {
		return &ErrInvalidPasswordOnSignup
	}
	return nil
}

func ( s *Server ) JoinChannelValidation ( joinchannel JoinChannel ) *ChannelError {
	if validation.Validate( joinchannel.Channel, validation.Required,
		validation.Match( IsUserName ), validation.Length(3, 16) ) != nil {
		return &ErrInvalidChannelName
	}
	if joinchannel.Comm != `SMS`  && joinchannel.Comm != `EMAIL` {
		return &ErrInvalidCommType
	}
	if validation.Validate( joinchannel.Alias, validation.Required, validation.Match( IsUserName ),
		validation.Length(3, 16) ) != nil {
		return &ErrInvalidAlias
	}
	return nil
}

func ( s *Server ) AddChannelValidation ( addchannel AddChannel ) *ChannelError {
	if validation.Validate(addchannel.Channel.Name, validation.Required, validation.Match(IsUserName),
		validation.Length(4, 50)) != nil {
		return &ErrInvalidChannelName
	}

	if &(addchannel.Channel.Preference) == nil {
		return &ErrEmptyChannelPreference
	}
	if len(addchannel.Channel.Restrictions) == 0 {
		return &ErrEmptyChannelRestrictions
	}
	if addchannel.Comm != `SMS`  && addchannel.Comm != `EMAIL` {
		return &ErrInvalidCommType
	}

	return nil
}

func ( s *Server ) AddCommValidation ( comm AddComm ) *ChannelError{
	if comm.CommType != `SMS`  && comm.CommType != `EMAIL` {
		return &ErrInvalidCommType
	}
	if comm.CommType == `SMS` && validation.Validate( comm.Comm, validation.Required, is.Digit,
		validation.Match( IsPhoneNumber ) ) != nil {
		return &ErrInvalidPhoneNumber
	}
	if comm.CommType == `EMAIL` && validation.Validate( comm.Comm, validation.Required, is.Email ) != nil {
		return &ErrInvalidEmail
	}
	return nil
}

func ( s *Server ) DeleteCommValidation ( comm DeleteComm ) *ChannelError{
	if validation.Validate( comm.Comm, validation.Required, is.Email ) != nil &&
		validation.Validate(comm.Comm, validation.Required, validation.Match(IsPhoneNumber) ) != nil {
		return &ErrInvalidCommType
	}
	return nil
}


func ( s *Server ) MessageValidation ( message SendMessage ) *ChannelError{
	if validation.Validate( message.Channel, validation.Required,
		validation.Match( IsUserName ), validation.Length(3, 16) ) != nil {
		return &ErrInvalidChannelName
	}
	return nil
}