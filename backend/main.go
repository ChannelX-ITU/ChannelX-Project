package main

import (
	"github.com/ChannelX-ITU/ChannelX-Project/backend/channel"
)

const (
	smtp = "smtp.gmail.com"
	port = 587
	username = "channel.x.itu@gmail.com"
	psswrd = "BestTeamEver"
)

func main()  {
	s := channel.Server{}
	s.Setup(smtp, port, username, psswrd)
	s.Run()
	s.Close()
}
