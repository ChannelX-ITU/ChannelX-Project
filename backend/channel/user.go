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

func (s *Server) GetAlias(userID int64, channelID int64) (alias string, er error) {
	get, err := s.dataBase.Prepare("SELECT A.val FROM CHANNEL_USER AS CU, ALIAS AS A WHERE CU.channel_id = ? AND CU.user_id = ? AND A.alias_id = CU.alias_id")
	if err != nil {
		return
	}

	err = get.QueryRow(channelID, userID).Scan(&alias)
	return
}

func (s *Server) UpdateAlias(userID int64, channelID int64, alias string) (err error) {
	set, err := s.dataBase.Prepare("UPDATE ALIAS SET val = ? WHERE alias_id = (SELECT DISTINCT FROM CHANNEL_USER WHERE channel_id = ? AND user_id = ?)")
	if err != nil {
		return
	}

	_, err = set.Exec(alias, channelID, userID)
	return
}

func (s *Server) DeleteUserIntervals(prefID int64) (err error) {
	del, err := s.dataBase.Prepare("DELETE FROM INTER WHERE preference_id = ?")

	if err != nil {
		return
	}

	defer del.Close()

	_, err = del.Exec(prefID)
	return
}

func (s *Server) UpdateUserPref(prefID int64, duration int, start int64) (err error) {
	upd, err := s.dataBase.Prepare("UPDATE PREFERENCE SET duration_days = ?, start_date = ? WHERE preference_id = ?")
	if err != nil {
		return
	}

	_, err = upd.Exec(duration, start, prefID)
	if err != nil {
	}
	return
}

func (s *Server) UpdateUser(userID int64, pref Preference) (err error) {
	prefID, err := s.GetPreference(userID)
	if err != nil {
		return
	}
	err = s.DeleteUserIntervals(prefID.prefID)
	if err != nil {
		return
	}
	err = s.AddInterval(pref.Intervals, prefID.prefID)
	if err != nil {
		return
	}
	err = s.UpdateChannelPref(prefID.prefID, pref.Duration, pref.StartDate)

	if err != nil {
		return
	}
	return
}