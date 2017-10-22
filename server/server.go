package server

import (
	"database/sql"
	"../channel"
	"net/http"
)

type Server struct {
	db		sql.DB
	mm		*channel.Mailman
	port 	string
}

func (s *Server) Setup(smtp string, port int, username string, psswrd string) {
	s.mm = &channel.Mailman{}
	s.mm.Setup(smtp, port, username, psswrd)
}

func (s *Server) Recieve(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Run() {
	s.mm.Run()
	http.HandleFunc("/", s.Recieve)
	http.ListenAndServe(":6969", nil)
}