package channel

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"encoding/json"
	"github.com/gorilla/sessions"
)
var store = sessions.NewCookieStore([]byte("bist-chinnil-ivir"))

type Server struct {
	dataBase		*sql.DB
	mailMan			*Mailman
	port 			string
}

func (s *Server) Setup(smtp string, port int, username string, psswrd string) {
	s.mailMan = &Mailman{}
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

func (s *Server) Run() {
	s.mailMan.Run()
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))

	router.HandleFunc("/", s.Receive)
	router.HandleFunc("/api/signUp", s.SignUp)
	router.HandleFunc("/api/login", s.Login)
	router.HandleFunc("/api/activate/{token}", s.ActivateToken)
	router.HandleFunc("/api/channels/{channel}", s.ServeChannel)
	router.HandleFunc("/api/channels/join", s.JoinChannelHandler)
	router.HandleFunc("/api/channels/add", s.AddChannelHandler)
	router.HandleFunc("/api/userinfo", s.ServeUserInfo)
	router.HandleFunc("/api/logout", s.Logout)
	router.HandleFunc("/api/channels", s.ServeChannels)
	router.HandleFunc("/api/comm/add", s.AddCommHandler)
	router.HandleFunc("/api/comm/remove", s.DeleteCommHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.HandleFunc("/{_:.*}", s.Receive)

	loggedHandler := handlers.CombinedLoggingHandler(os.Stdout, router)

	http.ListenAndServe(":6969", loggedHandler)
}


func (s *Server) Login(res http.ResponseWriter, req *http.Request) {
	// If method is GET serve an html login page
	if req.Method != "POST" {
		http.ServeFile(res, req, "static/templates/login.html")
		return
	}

	// Grab the username/password from the submitted post form
	decoder := json.NewDecoder(req.Body)
	var t Login
	err := decoder.Decode(&t)
	if err != nil {
		WriteError(res, ErrInternalServerError)
		return
	}
	defer req.Body.Close()
	username, password := t.Username, t.Password

	// Grab from the database
	var databaseUsername  string
	var databasePassword  string
	var databaseToken     string
	var userID			  int64

	// Search the database for the username provided
	// If it exists grab the password for validation
	err = s.dataBase.QueryRow("SELECT username, password FROM USERS WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
	// If not then redirect to the login page
	if err != nil {
		WriteError(res, ErrInvalidLoginCredentials)
		return
	}

	// Validate the password
	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	// If wrong password redirect to the login
	if err != nil {
		WriteError(res, ErrInvalidLoginCredentials)
		return
	}
	err = s.dataBase.QueryRow("SELECT token FROM TOKENS WHERE user_id=(SELECT user_id FROM USERS WHERE username=?)", username).Scan(&databaseToken)
	if err != sql.ErrNoRows {
		WriteError(res, ErrAccountNotActivated)
		return
	}

	err = s.dataBase.QueryRow("SELECT user_id FROM USERS WHERE username=?", username).Scan(&userID)
	if err != nil {
		WriteError(res, ErrInternalServerError)
		return
	}

	session, err := store.Get(req, "bist-sissin-ivir")
	if err != nil {
		WriteError(res, ErrInternalServerError)
		return
	}

	session.Values["user-id"] = userID
	session.Values["logged-in"] = true
	session.Save(req, res)
	WriteSuccess(res, "Login successful")
}

func (s *Server) SignUp(res http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		WriteError(res, ErrWrongMethod)
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t SignUp
	err := decoder.Decode(&t)
	if err != nil {
		WriteError(res, ErrInternalServerError)
		return
	}
	defer req.Body.Close()

	var user string
	username, email, password := t.Username, t.Email, t.Password

	err = s.dataBase.QueryRow("SELECT username FROM USERS WHERE username=?", username).Scan(&user)

	switch {
	// Username is available
	case err == sql.ErrNoRows:
		err := s.dataBase.QueryRow("SELECT C.val FROM COMM AS C, COMM_TYPE AS C_T WHERE C_T.val=? AND C.val=?", "EMAIL", email).Scan(&user)
		if err == nil {
			WriteError(res, ErrEmailTaken)
			return
		}
		if err != sql.ErrNoRows {
			WriteError(res, ErrInternalServerError)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			WriteError(res, ErrInternalServerError)
			return
		}

		_, err = s.dataBase.Exec("INSERT INTO USERS(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			WriteError(res, ErrInternalServerError)
			return
		}

		_, err = s.dataBase.Exec("INSERT INTO COMM(user_id, type_id, val) VALUES((SELECT U.user_id FROM USERS AS U WHERE U.username=?), (SELECT ct.type_id FROM COMM_TYPE AS ct WHERE ct.val=?), ?)", username, "EMAIL", email)
		if err != nil {
			WriteError(res, ErrInternalServerError)
			s.dataBase.Exec("DELETE FROM USERS WHERE username=?", username)
			return
		}

		_,err = s.dataBase.Exec("INSERT INTO PREFERENCE(user_id, start_date) VALUE ((SELECT U.user_id FROM USERS AS U WHERE U.username=?), 0)", username)
		if err != nil {
			WriteError(res, ErrInternalServerError)
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

		s.mailMan.Send(Message{email, "Activation", "To activate your account please click the link: http://localhost:6969/activate/" + u1.String()})
		WriteSuccess(res, "Activation mail is sent to the user's mail")
		return
	case err != nil:
		WriteError(res, ErrInternalServerError)
		return
	default:
		WriteError(res, ErrUsernameTaken)
	}
}

func (s *Server) Receive(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "static/index.html")
}


func (s *Server) ActivateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		WriteError(w, ErrWrongMethod)
		return
	}

	vars := mux.Vars(r)
	token := vars["token"]

	var databaseUser string
	err := s.dataBase.QueryRow("SELECT user_id FROM TOKENS WHERE token=?", token).Scan(&databaseUser)
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}
	s.dataBase.Exec("DELETE FROM TOKENS WHERE token=?", token)
	WriteSuccess(w, "Account is activated")
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}

	if !session.IsNew {
		cookie := &http.Cookie{
			Name:   "bist-sissin-ivir",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, cookie)
		WriteSuccess(w, "Logout successful")
	} else {
		// Set a new flash.
		WriteError(w, ErrNotLoggedIn)
		return
	}

}


func (s *Server) JoinChannelHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.IsNew {
		// Use the flash values.

		id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			decoder := json.NewDecoder(r.Body)
			var t JoinChannel
			err := decoder.Decode(&t)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}
			defer r.Body.Close()

			channelID, err := s.GetChannelID(t.Channel)

			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			if channelID == -1 {
				WriteError(w, ErrChannelNotExist)
				return
			}

			if ok, err := s.CheckUserInChannel(userId, channelID); err != nil {
				WriteError(w, ErrInternalServerError)
				return
			} else if ok {
				// Set a new flash.
				WriteError(w, ErrUserInChannel)
				return
			}

			commID, err := s.GetCommID(t.Comm)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			err = s.AddUserToChannel(channelID, userId, commID, false, t.Alias)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			WriteSuccess(w, "User is added to the channel")
		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		// Set a new flash.
		WriteError(w, ErrNotLoggedIn)
		return
	}
}

func (s *Server) AddChannelHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.IsNew {
		// Use the flash values.

		id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			decoder := json.NewDecoder(r.Body)
			var t AddChannel
			err := decoder.Decode(&t)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}
			defer r.Body.Close()

			channelID, _ := s.GetChannelID(t.Channel.Name)
			if channelID != -1 {
				WriteError(w, ErrChannelExist)
				return
			}

			err = s.AddChannel(t.Channel, userId, t.Comm)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}
			WriteSuccess(w, "Channel is created")

		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		// Set a new flash.
		WriteError(w, ErrNotLoggedIn)
		return
	}
}

func (s *Server) AddCommHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}

	if !session.IsNew {id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			decoder := json.NewDecoder(r.Body)
			var t AddComm
			err := decoder.Decode(&t)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}
			defer r.Body.Close()
			comm, commType := t.Comm, t.CommType

			err = s.AddComm(comm, commType, userId)
			if err == sql.ErrNoRows {
				WriteError(w, ErrCommIsTaken)
				return
			}

			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			WriteSuccess(w, "Communication method is added")
			return

		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		WriteError(w, ErrNotLoggedIn)
		return
	}
}

func (s *Server) DeleteCommHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}

	if !session.IsNew {id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			decoder := json.NewDecoder(r.Body)
			var t DeleteComm
			err := decoder.Decode(&t)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}
			defer r.Body.Close()
			comm := t.Comm

			err = s.DeleteComm(comm, userId)
			if err == sql.ErrNoRows {
				WriteError(w, ErrNoCommOfUser)
				return
			}

			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			WriteSuccess(w, "Communication method is deleted")
			return

		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		WriteError(w, ErrNotLoggedIn)
		return
	}
}


func (s *Server) ServeChannels (w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}

	if !session.IsNew {id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			ci, err := s.GetChannelInfos(userId)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			cib, err := json.Marshal(ci)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			w.Write(cib)
			return

		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		WriteError(w, ErrNotLoggedIn)
		return
	}
}

func (s *Server) ServerUserChannels(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}

	if !session.IsNew {
		// Use the flash values.

		id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			us, err := s.GetUser(userId)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			user, err := json.Marshal(us)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}


			w.Write(user)
		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		// Set a new flash.
		WriteError(w, ErrNotLoggedIn)
		return
	}
}

func (s *Server) ServeChannel(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		WriteError(w, ErrInternalServerError)
		return
	}

	if !session.IsNew {
		// Use the flash values.

		id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			vars := mux.Vars(r)
			channelName := vars["channel"]

			channelID, err := s.GetChannelID(channelName)

			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			if channelID == -1 {
				WriteError(w, ErrChannelNotExist)
				return
			}

			if ok, err := s.CheckUserInChannel(userId, channelID); err != nil {
				WriteError(w, ErrInternalServerError)
				return
			} else if !ok {
				// Set a new flash.
				WriteError(w, ErrUserNotInChannel)
				return
			}

			ch, err := s.GetChannel(channelID, userId)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			cha, err := json.Marshal(ch)
			if err != nil {
				WriteError(w, ErrInternalServerError)
				return
			}

			w.Write(cha)
		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		WriteError(w, ErrNotLoggedIn)
		return
	}
}

func (s *Server) ServeUserInfo(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		WriteError(w, ErrWrongMethod)
		return
	}

	session, err := store.Get(r, "bist-sissin-ivir")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.IsNew {
		// Use the flash values.

		id := session.Values["user-id"]
		if userId, ok := id.(int64); ok {
			us, err := s.GetUser(userId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			user, err := json.Marshal(us)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(user)
		} else {
			WriteError(w, ErrInternalServerError)
			return
		}
	} else {
		// Set a new flash.
		WriteError(w, ErrNotLoggedIn)
		return
	}
}


func (s *Server) Close() {
	s.dataBase.Close()
}
