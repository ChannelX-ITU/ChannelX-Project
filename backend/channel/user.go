package channel

import (
	"database/sql"
)

type User struct {
	Username		string 			`json:"username"`
	Channels		[]string		`json:"channels"`
	Preference		Preference		`json:"preferences"`
	Communications	[]Communication	`json:"communications"`
}

func (s *Server) GetUser(userID int64) (u User, err error) {
	u.Channels = make([]string, 0)
	u.Communications = make([]Communication, 0)

	u.Username, err = s.GetUsername(userID)
	if err != nil {
		return
	}
	u.Preference, err = s.GetPreference(userID)
	if err != nil {
		return
	}
	u.Communications, err = s.GetCommunication(userID)
	if err != nil {
		return
	}
	u.Channels, err = s.GetChannels(userID)
	if err != nil {
		return
	}

	return
}

func (s *Server) GetUsername(userID int64) (username string, err error) {
	get, err := s.dataBase.Prepare("SELECT username FROM USERS WHERE user_id = ?")
	if err != nil {
		return
	}

	err = get.QueryRow(userID).Scan(&username)
	return
}

func (s *Server) GetCommId(comm string, userID int64) (ID int64, err error) {
	get, err := s.dataBase.Prepare("SELECET C.comm_id FROM COMM AS C WHERE C.user_id = ? AND C.val = ?")
	if err != nil {
		return
	}

	err = get.QueryRow(userID, comm).Scan(&ID)
	if err == sql.ErrNoRows {
		return -1, nil
	}

	return
}

func (s *Server) AddComm(comm string, commType string, userID int64) (err error) {
	get, err := s.dataBase.Prepare("SELECT val FROM COMM WHERE val = ?")
	if err != nil {
		return
	}

	var st string

	err = get.QueryRow(comm).Scan(&st)
	if err != sql.ErrNoRows {
		err = sql.ErrTxDone
		return
	}

	get, err = s.dataBase.Prepare("INSERT INTO COMM (user_id, type_id, val) VALUES (?, ?, ?)")
	if err != nil {
		return
	}

	t, err := s.GetCommType(commType)
	if err != nil {
		return
	}

	_, err = get.Exec(userID, t, comm)

	defer get.Close()
	return
}

func (s *Server) DeleteComm(comm string, userID int64) (err error) {
	del, err := s.dataBase.Prepare("DELETE FROM COMM WHERE val = ? AND user_id = ?")
	if err != nil {
		return
	}

	defer del.Close()

	res, err := del.Exec(comm, userID)
	if err != nil {
		return sql.ErrTxDone
	}

	if res != nil {
		if n, _ := res.RowsAffected(); n == 0 {
			return sql.ErrNoRows
		}
	}


	return
}