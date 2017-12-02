package channel

import "fmt"

type Channel struct {
	Name			string		`json:"name"`
	IsOwner			int			`json:"is_owner,omitempty"`
	Preference		Preference	`json:"preference"`
	Restrictions	Restriction	`json:"restrictions"`
	Users			[]string	`json:"users,omitempty"`
}

func (s *Server) GetChannels(userID int) (inf []string, err error) {
	inf = make([]string, 0)
	get, err := s.dataBase.Prepare("SELECT DISTINCT C.name FROM USERS AS U, CHANNEL AS C, CHANNEL_USER AS CU WHERE C.channel_id = CU.channel_id AND CU.user_id = ?")
	if err != nil {
		return
	}

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

func (s *Server) AddChannel(channel Channel, userID int) (err error) {
	set, err := s.dataBase.Prepare("INSERT INTO CHANNEL(name) VALUES(?)")
	if err != nil {
		return
	}

	res, err := set.Exec(channel.Name)
	if err != nil {
		return
	}

	channelID,err := res.LastInsertId()
	if err != nil {
		return
	}

	fmt.Println("Channel inserted has ID:", channelID)

	err = s.AddPreference(channel.Preference, int(channelID))
	if err != nil {
		return
	}

	return
}