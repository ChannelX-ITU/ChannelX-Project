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

	/*
	defer rows.Close()

	for rows.Next() {

		// Scan the value to []byte
		err = rows.Scan(&u.Username, &channelName, &u.Preference.StartDate, &u.Preference.Duration, &com.Value, &com.Type)
		fmt.Println(u.Username)

		if err != nil {
			fmt.Print(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
			continue
		}

		sort.Strings(u.Channels)

		if sort.SearchStrings(u.Channels, channelName) >= len(u.Channels) {
			u.Channels = append(u.Channels, channelName)
		}

		if !ComInSlice(u.Communications, com) {
			u.Communications = append(u.Communications, com)
		}
	}
	*/

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