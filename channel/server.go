package channel

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"encoding/json"
)

type Server struct {
	dataBase		*sql.DB
	mailMan			*Mailman
	port 			string
}

func (s *Server) Setup(smtp string, port int, username string, psswrd string) {
	s.mailMan = &Mailman{}
	s.mailMan.Setup(smtp, port, username, psswrd)
	db, err := sql.Open("mysql", "root:35792030@/ChannelX")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	s.dataBase = db
}

func (s *Server) login(res http.ResponseWriter, req *http.Request) {
	// If method is GET serve an html login page
	if req.Method != "POST" {
		http.ServeFile(res, req, "static/templates/login.html")
		return
	}

	// Grab the username/password from the submitted post form
	username := req.FormValue("username")
	password := req.FormValue("password")

	// Grab from the database
	var databaseUsername  string
	var databasePassword  string
	var databaseToken     string
	var userID			  int

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
	err = s.dataBase.QueryRow("SELECT token FROM TOKENS WHERE user_id=(SELECT user_id FROM USERS WHERE username=?)", username).Scan(&databaseToken)
	if err != sql.ErrNoRows {
		res.Write([]byte("Please activate your account!"))
		return
	}

	err = s.dataBase.QueryRow("SELECT user_id FROM USERS WHERE username=?", username).Scan(&userID)
	if err != nil {
		fmt.Println(err)
		res.Write([]byte("Server error"))
		return
	}

	// If the login succeeded
	us, err := s.GetUser(userID)
	if err != nil {
		fmt.Println(err)
		res.Write([]byte("Server error"))
		return
	}

	user, err := json.Marshal(us)
	if err != nil {
		fmt.Println(err)
		res.Write([]byte("Server error"))
		return
	}
	fmt.Println(s.AddChannel(Channel{"IdsaNdSfooC32NM", 1, Preference{2131231321, 323, []Interval{Interval{21323, 2132}}}, []Restriction{Restriction{19, "EMAIL", "@itu.edu.tr", "=", "END"}}, make([]string, 0)}, 32, "cicekhu@gmail.com").Error())

	res.Write(user)
}

func (s *Server) SingupPage(res http.ResponseWriter, req *http.Request) {

	// Serve signup.html to get requests to /signup
	if req.Method != "POST" {
		http.ServeFile(res, req, "static/templates/signup.html")
		return
	}
	email	 := req.FormValue("email")
	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	err := s.dataBase.QueryRow("SELECT username FROM USERS WHERE username=?", username).Scan(&user)

	switch {
	// Username is available
	case err == sql.ErrNoRows:
		err := s.dataBase.QueryRow("SELECT C.val FROM COMM AS C, COMM_TYPE AS C_T WHERE C_T.val=? AND C.val=?", "EMAIL", email).Scan(&user)
		if err == nil {
			res.Write([]byte("E-Mail: " + email + " is already taken!"))
			return
		}
		if err != sql.ErrNoRows {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			http.Error(res, "4Server error, unable to create your account.", 500)
			return
		}

		_, err = s.dataBase.Exec("INSERT INTO USERS(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(res, "3Server error, unable to create your account.", 500)
			return
		}

		_, err = s.dataBase.Exec("INSERT INTO COMM(user_id, type_id, val) VALUES((SELECT U.user_id FROM USERS AS U WHERE U.username=?), (SELECT ct.type_id FROM COMM_TYPE AS ct WHERE ct.val=?), ?)", username, "EMAIL", email)
		if err != nil {
			http.Error(res, "2Server error, unable to create your account.", 500)
			s.dataBase.Exec("DELETE FROM USERS WHERE username=?", username)
			return
		}

		_,err = s.dataBase.Exec("INSERT INTO PREFERENCE(user_id, start_date) VALUE ((SELECT U.user_id FROM USERS AS U WHERE U.username=?), 0)", username)
		if err != nil {
			http.Error(res, "1Server error, unable to create your account.", 500)
			s.dataBase.Exec("DELETE FROM COMM WHERE val=?", email)
			s.dataBase.Exec("DELETE FROM USERS WHERE username=?", username)
			return
		}

		u1 := uuid.NewV4()

		_,err = s.dataBase.Exec("INSERT INTO TOKENS(user_id, token) VALUE ((SELECT U.user_id FROM USERS AS U WHERE U.username=?), ? )", username, u1)
		if err != nil {
			http.Error(res, err.Error(), 500)
			s.dataBase.Exec( "DELETE FROM PREFERENCE WHERE user_id=(SELECT  U.user_id FROM USERS AS U WHERE U.username=?)", username )
			s.dataBase.Exec("DELETE FROM COMM WHERE val=?", email)
			s.dataBase.Exec("DELETE FROM USERS WHERE username=?", username)
			return
		}

		s.mailMan.Send(Message{email, "Activation", "To activate your account please click the link: localhost:6969/activate/" + u1.String()})
		res.Write([]byte("Activation mail is sent to " + email))
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		res.Write([]byte("Username: " + user + " is already taken!"))
	}
}

func (s *Server) Recieve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selam"))
}

func (s *Server) Run() {
	s.mailMan.Run()
	router := mux.NewRouter()
	router.HandleFunc("/", s.Recieve)
	router.HandleFunc("/signup", s.SingupPage)
	router.HandleFunc("/submitsignup", s.SubmitSignUp)
	router.HandleFunc("/login", s.login)
	router.HandleFunc("/activate/{token}", s.ActivateToken)

	http.ListenAndServe(":6969", router)
}

func (s *Server) ActivateToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	var databaseUser string
	err := s.dataBase.QueryRow("SELECT user_id FROM TOKENS WHERE token=?", token).Scan(&databaseUser)
	if err != nil {
		w.Write([]byte("This activation token is not valid!"))
		return
	}
	s.dataBase.Exec("DELETE FROM TOKENS WHERE token=?", token)
	w.Write([]byte("Your account has been activated!"))
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

func (s *Server) Close() {
	s.dataBase.Close()
}