package main

import (
	"github.com/ChannelX-ITU/ChannelX-Project/backend/server"
)

const (
	smtp = "smtp.gmail.com"
	port = 465
	username = "channel.x.itu@gmail.com"
	psswrd = "BestTeamEver"
)

func main()  {
	s := server.Server{}
	s.Setup(smtp, port, username, psswrd)
	s.Run()
}
