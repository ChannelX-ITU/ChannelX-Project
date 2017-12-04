package channel

import (
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

const (
	timeOut = 30 * time.Second
)

type Mailman struct {
	mc		chan Message
	smtp		string
	port		int
	username	string
	password	string
	dial 		*gomail.Dialer
}

func (m *Mailman) Setup(smtp string, port int, username string, password string) {
	m.mc = make(chan Message, 100)
	m.smtp = smtp
	m.port = port
	m.username = username
	m.password = password
	m.dial = gomail.NewDialer(m.smtp, m.port, m.username, m.password)

	
}

func (m *Mailman) Run()  {

	go func() {

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case msg, ok := <-m.mc:
				if !ok {
					return
				}
				if !open {
					if s, err = m.dial.Dial(); err != nil {
						panic(err)
					}
					open = true
				}

				ma := gomail.NewMessage()
				ma.SetHeader("From", m.username)
				ma.SetHeader("To", msg.To)
				ma.SetHeader("Subject", msg.Sub)
				ma.SetBody("text/html", msg.Msg)

				if err := gomail.Send(s, ma); err != nil {
					log.Print(err)
				}
			case <-time.After(timeOut):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()
}

func (m *Mailman) Send(msg Message) {
	m.mc <- msg
}

func (m *Mailman) Stop() {
	close(m.mc)
}
