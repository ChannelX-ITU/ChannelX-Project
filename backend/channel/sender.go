package channel

type Sender interface {
	Send(Message)
}