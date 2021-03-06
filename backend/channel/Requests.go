package channel

type Login struct {
	Username	string `json:"username"`
	Password	string `json:"password"`
}

type SignUp struct {
	Email		string `json:"email"`
	Username	string `json:"username"`
	Password	string `json:"password"`
}

type JoinChannel struct {
	Channel		string	`json:"channel"`
	Comm		string	`json:"comm"`
	Alias		string	`json:"alias,omitempty"`
}

type AddChannel struct {
	Channel		Channel	`json:"channel"`
	Comm		string	`json:"comm"`
	Alias		string	`json:"alias,omitempty"`
}

type AddComm struct {
	Comm		string	`json:"value"`
	CommType	string	`json:"comm_type"`
}

type DeleteComm struct {
	Comm		string	`json:"value"`
}

type SendMessage struct {
	Channel		string	`json:"channel"`
	Subject		string
	Message		string	`json:"message"`
}

type LeaveChannel struct {
	Channel		string	`json:"channel"`
}

type ChannelWrapper struct {
	Channel Channel 	`json:"channel"`
	Comm	Communication	`json:"comm"`
	Alias 		string 	`json:"alias"`
}
