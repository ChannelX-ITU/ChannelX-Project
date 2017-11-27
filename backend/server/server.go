package server

import (
	"github.com/ChannelX-ITU/ChannelX-Project/backend/channel"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Server struct {
	dataBase		*sql.DB
	mailMan			*channel.Mailman
	port 			string
}

func (s *Server) Setup(smtp string, port int, username string, psswrd string) {
	s.mailMan = &channel.Mailman{}
	s.mailMan.Setup(smtp, port, username, psswrd)
	db, err := sql.Open("mysql", "root:35792030@tcp(my.sql:3306)/ChannelX")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	s.dataBase = db
}

func (s *Server) Login(res http.ResponseWriter, req *http.Request) {
	// If method is GET serve an html login page
	if req.Method != "POST" {
		http.Redirect(res, req, "/login", 301)
		return
	}

	// Grab the username/password from the submitted post form
	username := req.FormValue("username")
	password := req.FormValue("password")

	// Grab from the database
	var databaseUsername  string
	var databasePassword  string

	// Search the database for the username provided
	// If it exists grab the password for validation
	err := s.dataBase.QueryRow("SELECT username, password FROM USERS WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
	// If not then redirect to the login page
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	// Validate the password
	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	// If wrong password redirect to the login
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	// If the login succeeded
	res.Write([]byte("Hello " + databaseUsername))
}

func (s *Server) SignupPage(res http.ResponseWriter, req *http.Request) {

	// Serve signup.html to get requests to /signup
	if req.Method != "POST" {
		http.Redirect(res, req, "/register", 301)
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	err := s.dataBase.QueryRow("SELECT username FROM USERS WHERE username=?", username).Scan(&user)

	switch {
	// Username is available
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account. 1: " + err.Error(), 500)
			return
		}

		_, err = s.dataBase.Exec("INSERT INTO USERS(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(res, "Server error, unable to create your account. 2: " + err.Error(), 500)
			return
		}

		res.Write([]byte("User created!"))
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account. 3: " + err.Error(), 500)

		return
	default:
		res.Write([]byte("Username: " + user + " is already taken!"))
		time.Sleep(3 * time.Second)
	}
}

func (s *Server) Recieve(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "static/index.html")
}

func (s *Server) Run() {
	s.mailMan.Run()
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	router.HandleFunc("/api/login", s.Login)
	router.HandleFunc("/api/signup", s.SignupPage)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.HandleFunc("/{_:.*}", s.Recieve)

	// router.HandleFunc("/submitsignup", s.SubmitSignUp)

	loggedHandler := handlers.CombinedLoggingHandler(os.Stdout, router)
	http.ListenAndServe(":6969", loggedHandler)
}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/signup.html")
}

func (s *Server) SubmitSignUp(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("A form is submitted!")
	r.ParseForm()
	for key, val :=  range r.Form {
		fmt.Println("Key:", key, "Value:", val)
	}
	w.Write([]byte("Danke!"))
}
