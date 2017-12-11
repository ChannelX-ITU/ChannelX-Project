package channel

import (
	"regexp"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var IsUserName 		= regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)
var IsPhoneNumber	= regexp.MustCompile( `(0|\\+90|090|90)[0-9]+$`)

//tested
func ( s *Server ) LoginValidation ( login Login ) *ChannelError {
	if validation.Validate( login.Username, validation.Required,
		validation.Match( IsUserName ), validation.Length(4, 30) ) != nil {
		return &ErrInvalidLoginCredentials
	}
	return nil
}
//tested
func ( s *Server ) SignupValidation ( signup SignUp ) *ChannelError {
	if validation.Validate(signup.Email, validation.Required, is.Email) != nil {
		return &ErrInvalidEmail
	}
	if validation.Validate( signup.Username, validation.Required,
		validation.Match( IsUserName ), validation.Length(4, 30) ) != nil {
		return &ErrInvalidUsernameOnSignup
	}
	if validation.Validate( signup.Password, validation.Required, validation.Length(6, 50) ) != nil {
		return &ErrInvalidPasswordOnSignup
	}
	return nil
}
//tested
func ( s *Server ) JoinChannelValidation ( joinchannel JoinChannel ) *ChannelError {
	if validation.Validate( joinchannel.Channel, validation.Required,
		validation.Match( IsUserName ), validation.Length(4, 30) ) != nil {
		return &ErrInvalidChannelName
	}
	if validation.Validate( joinchannel.Comm, validation.Required, validation.Match( IsPhoneNumber ),
		validation.Length(7, 15) ) != nil  &&
		validation.Validate(joinchannel.Comm, validation.Required, is.Email) != nil {
		return &ErrInvalidCommType
	}
	return nil
}
//tested
func ( s *Server ) AddChannelValidation ( addchannel AddChannel ) *ChannelError {
	if validation.Validate(addchannel.Channel.Name, validation.Required, validation.Match(IsUserName),
		validation.Length(4, 30)) != nil {
		return &ErrInvalidChannelName
	}

	if &(addchannel.Channel.Preference) == nil {
		return &ErrEmptyChannelPreference
	}
	if validation.Validate(addchannel.Comm, validation.Required, validation.Match(IsPhoneNumber), validation.Length(7, 15),) != nil &&
		validation.Validate(addchannel.Comm, validation.Required, is.Email,) != nil {
		return &ErrInvalidCommType
	}

	return nil
}
//tested
func ( s *Server ) AddCommValidation ( comm AddComm ) *ChannelError{
	if comm.CommType != `SMS`  && comm.CommType != `EMAIL` {
		return &ErrInvalidCommType
	}
	if comm.CommType == `SMS` && validation.Validate( comm.Comm, validation.Required,
		validation.Match( IsPhoneNumber ), validation.Length(7, 15) ) != nil {
		return &ErrInvalidPhoneNumber
	}
	if comm.CommType == `EMAIL` && validation.Validate( comm.Comm, validation.Required, is.Email ) != nil {
		return &ErrInvalidEmail
	}
	return nil
}
//tested
func ( s *Server ) DeleteCommValidation ( comm DeleteComm ) *ChannelError{
	if validation.Validate( comm.Comm, validation.Required, is.Email ) != nil &&
		validation.Validate(comm.Comm, validation.Required, validation.Match(IsPhoneNumber), validation.Length(7, 15), ) != nil {
		return &ErrInvalidCommType
	}
	return nil
}
//tested
func ( s *Server ) MessageValidation ( message SendMessage ) *ChannelError{
	if validation.Validate( message.Channel, validation.Required,
		validation.Match( IsUserName ), validation.Length(4, 30) ) != nil {
		return &ErrInvalidChannelName
	}
	return nil
}
//tested
func ( s *Server ) LeaveChannelValidation ( channel LeaveChannel ) *ChannelError{
	if validation.Validate( channel.Channel, validation.Required,
		validation.Match( IsUserName ), validation.Length(4, 30) ) != nil {
		return &ErrInvalidChannelName
	}
	return nil
}
