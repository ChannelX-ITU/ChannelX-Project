package channel

import (
	"database/sql"
)

type Channel struct {
	Name			string			`json:"name"`
	IsOwner			bool				`json:"is_owner,omitempty"`
	Preference		Preference		`json:"preference"`
	Restrictions	[]Restriction	`json:"restrictions"`
	Users			[]string		`json:"users,omitempty"`
}

func (s *Server) GetChannels(userID int64) (inf []string, err error) {
	inf = make([]string, 0)
	get, err := s.dataBase.Prepare("SELECT DISTINCT C.name FROM USERS AS U, CHANNEL AS C, CHANNEL_USER AS CU WHERE C.channel_id = CU.channel_id AND CU.user_id = ?")
	if err != nil {
		return
	}

	defer get.Close()

	rows, err := get.Query(userID)
	if rows == nil || err != nil {
		return
	}

	defer rows.Close()

	var i string

	for rows.Next() {
		rows.Scan(&i)
		inf = append(inf, i)
	}

	return
}

func (s *Server) AddChannel(channel Channel, userID int64, comm string) (err error) {
	set, err := s.dataBase.Prepare("INSERT INTO CHANNEL(name) VALUES(?)")
	if err != nil {
		return
	}

	defer set.Close()

	res, err := set.Exec(channel.Name)
	if err != nil {
		return
	}

	channelID,err := res.LastInsertId()
	if err != nil {
		return
	}

	prefID, err := s.AddPreference(channel.Preference, channelID)
	if err != nil {
		return
	}

	err = s.AddRestriction(channel.Restrictions, prefID)
	if err != nil {
		return
	}

	commID, err := s.GetCommID(comm)
	if err != nil {
		return
	}

	err = s.AddUserToChannel(int64(channelID), int64(userID), int(commID), true, "")

	return
}

func (s *Server) AddUserToChannel(channelID int64, userID int64, commID int, isOwner bool, alias string) (err error) {


	setAlias, err := s.dataBase.Prepare("INSERT INTO ALIAS(val, is_user_defined) VALUES(?, ?)")
	if err != nil {
		return
	}


	defer setAlias.Close()

	var res sql.Result

	var userDefined bool
	if alias != "" {
		userDefined = true
		res, err = setAlias.Exec(alias, userDefined)
		if err != nil {
			return
		}
	} else {
		userDefined = false
		res, err = setAlias.Exec("Boring Panda", userDefined)
		if err != nil {
			return
		}
	}


	aliasID, err := res.LastInsertId()
	if err != nil {
		return
	}


	set, err := s.dataBase.Prepare("INSERT INTO CHANNEL_USER(channel_id, comm_id, alias_id, is_owner, user_id) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return
	}


	defer set.Close()

	_, err = set.Exec(channelID, commID, aliasID, isOwner, userID)

	return
}

func (s *Server) GetCommID(comm string) (commID int64, err error) {
	get, err := s.dataBase.Prepare("SELECT comm_id FROM COMM WHERE val = ?")
	if err != nil {
		return
	}

	defer get.Close()

	err = get.QueryRow(comm).Scan(&commID)
	return
}

func (s *Server) GetChannel(channelID int64, userID int64) (ch Channel, err error) {
	ch.Name, err = s.GetChannelName(channelID)
	if err != nil {
		return
	}

	var ownerID int64
	ch.Users, ownerID, err = s.GetChannelUsers(channelID)
	if err != nil {
		return
	}

	if ownerID == userID {
		ch.IsOwner = true
	}

	ch.Preference, err = s.GetPreferenceForChannel(channelID)
	if err != nil {
		return
	}

	ch.Restrictions, err = s.GetRestrictions(ch.Preference.prefID)
	if err != nil {
		return
	}

	return
}

func (s *Server) GetChannelName(channelID int64) (name string, err error) {
	get, err := s.dataBase.Prepare("SELECT name FROM CHANNEL WHERE channel_id = ?")
	if err != nil {
		return
	}

	defer get.Close()

	err = get.QueryRow(channelID).Scan(&name)
	return
}

func (s *Server) GetChannelUsers(channelID int64) (users []string, owner int64, err error) {
	get, err := s.dataBase.Prepare("SELECT CU.is_owner, A.val, CU.user_id FROM ALIAS AS A, CHANNEL_USER AS CU, CHANNEL AS C WHERE A.alias_id = CU.alias_id AND CU.channel_id = C.channel_id AND C.channel_id = ?")

	rows, err := get.Query(channelID)
	if rows == nil || err != nil {
		return
	}

	defer rows.Close()

	users = make([]string, 0)
	var hold string
	var isOwner bool
	var ownerID int64
	for rows.Next() {
		rows.Scan(&isOwner, &hold,&ownerID)
		users = append(users, hold)
		if isOwner {
			owner = ownerID
		}
	}

	return
}

func (s *Server) GetChannelID(name string) (ID int64, err error) {
	get, err := s.dataBase.Prepare("SELECT C.channel_id FROM CHANNEL AS C WHERE C.name = ?")
	if err != nil {
		return
	}

	err = get.QueryRow(name).Scan(&ID)
	if err == sql.ErrNoRows {
		return -1, nil
	}

	return
}

func (s *Server) CheckUserInChannel(userID int64, channelID int64) (ok bool, err error) {
	get, err := s.dataBase.Prepare("SELECT channel_id FROM CHANNEL_USER WHERE user_id = ? AND channel_id = ?")
	if err != nil {
		return
	}

	var ID int64

	err = get.QueryRow(userID, channelID).Scan(&ID)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, err
}