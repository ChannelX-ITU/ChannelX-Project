package channel

import (
	"database/sql"
	"time"
	"github.com/satori/go.uuid"
	"fmt"
)

type Channel struct {
	Name			string			`json:"name"`
	IsOwner			bool			`json:"is_owner,omitempty"`
	Preference		Preference		`json:"preference"`
	Restrictions	[]Restriction	`json:"restrictions"`
	Users			[]string		`json:"users,omitempty"`
}

type ChannelInfo struct {
	Name	string		`json:"name"`
	Count	int			`json:"user_count"`
	Comm	string		`json:"comm"`
}

type ChannelsInfo struct {
	Owned	[]ChannelInfo	`json:"owned"`
	Subbed	[]ChannelInfo	`json:"subbed"`
}

func (s *Server) GetChannelInfo(userID int64, channelName string) (channelInfo ChannelInfo, isOwner bool, err error) {
	channelID,err := s.GetChannelID(channelName)
	if err != nil {
		return
	}

	ch, err := s.GetChannel(channelID, userID)
	if err != nil {
		return
	}

	comm, err := s.GetCommOfUserInChannel(channelID, userID)
	if err != nil {
		return
	}

	isOwner, err = s.GetIsUserOwner(channelID, userID)
	if err != nil {
		return
	}

	channelInfo.Name = ch.Name
	channelInfo.Count = len(ch.Users)
	channelInfo.Comm = comm
	return
}

func (s *Server) GetIsUserOwner(chanID int64, userID int64) (ok bool, err error) {
	get, err := s.dataBase.Prepare("SELECT DISTINCT CU.is_owner FROM CHANNEL_USER AS CU WHERE CU.user_id = ? AND CU.channel_id =?")
	if err != nil {
		return
	}

	defer get.Close()

	err = get.QueryRow(userID, chanID).Scan(&ok)
	return
}

func (s *Server) GetChannelInfos(userID int64) (ci ChannelsInfo, err error) {
	ci.Owned = make([]ChannelInfo, 0)
	ci.Subbed = make([]ChannelInfo, 0)

	arr, err := s.GetChannels(userID)
	if err != nil {
		return
	}

	for _, i := range arr {
		cu, isOwner, err := s.GetChannelInfo(userID, i)
		if err != nil {
			continue
		}

		if isOwner {
			ci.Owned = append(ci.Owned, cu)
		} else {
			ci.Subbed = append(ci.Subbed, cu)
		}
	}

	err = nil
	return
}

func (s *Server) GetCommOfUserInChannel(chanID int64, userID int64) (comm string, err error) {
	get, err := s.dataBase.Prepare("SELECT DISTINCT C.val FROM CHANNEL_USER AS CU, COMM AS C WHERE CU.user_id = ? AND CU.channel_id = ? AND C.comm_id = CU.comm_id")
	if err != nil {
		return
	}

	defer get.Close()

	err = get.QueryRow(userID, chanID).Scan(&comm)
	return
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

	err = s.AddUserToChannel(int64(channelID), int64(userID), commID, true, "")

	return
}

func (s *Server) AddUserToChannel(channelID int64, userID int64, commID int64, isOwner bool, alias string) (err error) {


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


	set, err := s.dataBase.Prepare("INSERT INTO CHANNEL_USER(channel_id, comm_id, alias_id, is_owner, user_id, token) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}


	defer set.Close()

	_, err = set.Exec(channelID, commID, aliasID, isOwner, userID, uuid.NewV4())

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

func (s *Server) GetAllCommInChannel(channelID int64) (comms []Communication, err error) {
	comms = make([]Communication, 0)
	get, err := s.dataBase.Prepare("SELECT C.val, CT.val FROM CHANNEL_USER AS CU, COMM AS C, COMM_TYPE AS CT WHERE CU.channel_id = ? AND CU.comm_id = C.comm_id AND C.type_id = CT.type_id")
	if err != nil {
		return
	}

	defer get.Close()

	rows, err := get.Query(channelID)
	if err != nil || rows == nil {
		return
	}

	for rows.Next() {
		i := Communication{}
		err = rows.Scan(&i.Value, &i.Type)
		if err != nil {
			return
		}

		comms = append(comms, i)
	}

	err = nil
	return
}

func (s *Server) GetOwnerCommInChannel(channelID int64) (comm Communication, err error) {
	get, err := s.dataBase.Prepare("SELECT C.val, CT.val FROM CHANNEL_USER AS CU, COMM AS C, COMM_TYPE AS CT WHERE CU.channel_id = ? AND CU.comm_id = C.comm_id AND C.type_id = CT.type_id AND CU.is_owner")
	if err != nil {
		return
	}
	defer get.Close()

	err = get.QueryRow(channelID).Scan(&comm.Value, &comm.Type)
	return
}

func (s *Server) CheckTimeForSend(channelID int64) (ok bool, err error) {
	ok = false
	var(
		prefId	int64
		stDate	int64
		drDays	int
	)
	var intervals[] Interval

	err = s.dataBase.QueryRow("SELECT preference_id FROM PREFERENCE WHERE channeL_id=?", channelID).Scan(&prefId)
	if err != nil {
		return
	}

	err = s.dataBase.QueryRow("SELECT start_date, duration_days FROM PREFERENCE WHERE preference_id=?", prefId).Scan(&stDate, &drDays)
	if err != nil {
		return
	}

	err = s.dataBase.QueryRow("SELECT start_date, duration_days FROM PREFERENCE WHERE preference_id=?", prefId).Scan(&stDate, &drDays)
	if err != nil {
		return
	}

	intervals, err = s.GetInterval(prefId)
	if err != nil{
		return
	}

	now := time.Now().UTC()
	nowMSeconds := now.UnixNano() / 1000000

	endDate := stDate + int64(drDays)*(8.64e+7)
	//buraya kadar sorun yok

	//var ownerUserID int64

	//day control
	if nowMSeconds < stDate || nowMSeconds > endDate {
		if nowMSeconds > endDate {
			if drDays != 0 {
				/*
				err = s.dataBase.QueryRow("SELECT CU.user_id FROM CHANNEL, CHANNEL_USER AS CU WHERE CHANNEL.channel_id=CU.channel_id AND CHANNEL.channel_id=? AND CU.is_owner=1", channelID).Scan(&ownerUserID)
				if err != nil {
					return
				}
				s.DeleteUserFromChannel(channelID, ownerUserID, true)
				*/
				return
			}
		} else{// burada patliyor, cunku bakiyor ki kanal daha acilmamis
			fmt.Println("at\n\n\n\n\n\n\n")
			return
		}
	}
	if len(intervals) == 0 {
		return true, nil
	}
	var nowValue int
	for _, j := range intervals {
		if j.Start / 1440 == ( (int(now.Weekday()) + 6) % 7 ) {	//if its the day message is allowed
			nowValue = 1440 * ( j.Start / 1440 )  + now.Hour()*60 + now.Minute()
			if j.Start <= nowValue && nowValue <= ( j.Start + j.Duration ) {
				return true, nil
			}
		}
	}

	return false, err
}

func (s *Server) DeleteUserFromChannel( channelID int64, userID int64, isOwner bool) (err error) {
	var data string

	if isOwner {	//owner ise
		err = s.dataBase.QueryRow("SELECT preference_id FROM PREFERENCE WHERE channel_id=?", channelID).Scan(&data)//pref data is needed for restriction and interval infos
		if err != nil {
			return
		}
		_, err = s.dataBase.Exec("DELETE FROM INTER WHERE preference_id=?", data)//interval is deleted
		if err != nil {
			return
		}
		_, err = s.dataBase.Exec("DELETE FROM RESTRICTION WHERE preference_id=?", data)//restrictions are deleted
		if err != nil {

			return
		}

		_, err = s.dataBase.Exec("DELETE FROM PREFERENCE WHERE preference_id=?", data)//preference is deleted
		if err != nil {
			return
		}

		arr := make([]int64, 0)

		row, err := s.dataBase.Query("SELECT alias_id FROM CHANNEL_USER WHERE channel_id=?", channelID)
		if err != nil || row == nil {
			return err
		}

		var i int64

		for row.Next() {
			err = row.Scan(&i)
			if err != nil {
				continue
			}

			arr = append(arr, i)
		}

		_, err = s.dataBase.Exec("DELETE FROM CHANNEL_USER WHERE channel_id=?", channelID)//every channel users are deleted
		if err != nil {
			return err
		}

		for _, j := range arr {
			_, err = s.dataBase.Exec("DELETE FROM ALIAS WHERE alias_id=?", j)//every single alias' for each user in channel are deleted
			if err != nil {
				continue
			}
		}

		_, err = s.dataBase.Exec("DELETE FROM CHANNEL WHERE channel_id=?", channelID)//every channel users are deleted
		if err != nil {
			return err
		}

		return nil
	}else {
		err = s.dataBase.QueryRow("SELECT alias_id FROM CHANNEL_USER WHERE channel_id=? AND user_id=?", channelID, userID).Scan(&data)
		_, err = s.dataBase.Exec("DELETE FROM CHANNEL_USER WHERE channel_id=? AND user_id=?", channelID, userID)//channel_user is deleted
		if err != nil {
			return
		}

		_, err = s.dataBase.Exec("DELETE FROM ALIAS WHERE alias_id=?", data)//alias is deleted
		if err != nil {
			return
		}
	}
	return nil
}

func (s *Server) GetChannelUserFromToken(token string) (channelID int64, userID int64, err error) {

	get, err := s.dataBase.Prepare("SELECT CU.user_id, CU.channel_id FROM CHANNEL_USER AS CU WHERE CU.token = ?")
	if err != nil {
		return
	}

	defer get.Close()

	err = get.QueryRow(token).Scan(&userID, &channelID)
	return
}

func (s *Server) GetChannelUserToken(channelID int64, userID int64) (token string, err error) {
	get, err := s.dataBase.Prepare("SELECT DISTINCT CU.token FROM CHANNEL_USER AS CU WHERE CU.user_id = ? AND CU.channel_id =?")
	if err != nil {
		return
	}

	defer get.Close()

	err = get.QueryRow(userID, channelID).Scan(&token)

	return
}

func (s *Server) DeleteChannelIntervals(prefID int64) (err error) {
	del, err := s.dataBase.Prepare("DELETE FROM INTER WHERE preference_id = ?")

	if err != nil {
		return
	}

	defer del.Close()

	_, err = del.Exec(prefID)
	return
}

func (s *Server) DeleteChannelRestrictions(prefID int64) (err error) {
	del, err := s.dataBase.Prepare("DELETE FROM RESTRICTION WHERE preference_id = ?")

	if err != nil {
		return
	}

	defer del.Close()

	_, err = del.Exec(prefID)
	return
}

func (s *Server) UpdateChannelPref(prefID int64, duration int, start int64) (err error) {
	upd, err := s.dataBase.Prepare("UPDATE PREFERENCE SET duration_days = ?, start_date = ? WHERE preference_id = ?")
	if err != nil {
		return
	}

	_, err = upd.Exec(duration, start, prefID)
	if err != nil {
	}
	return
}

func (s *Server) UpdateChannel(userID int64, cha Channel, comm string, alias string) (err error) {
	channelID, err := s.GetChannelID(cha.Name)
	if err != nil {
		return
	}

	prefID, err := s.GetPreferenceForChannel(channelID)
	if err != nil {
		return
	}

	s.DeleteChannelIntervals(prefID.prefID)
	s.DeleteChannelRestrictions(prefID.prefID)
	s.UpdateChannelPref(prefID.prefID, cha.Preference.Duration, cha.Preference.StartDate)
	s.AddInterval(cha.Preference.Intervals, prefID.prefID)
	s.AddRestriction(cha.Restrictions, prefID.prefID)
	s.UpdateAlias(userID, channelID, alias)
	s.UpdateCommInChannel(userID, channelID, comm)
	return
}

func (s *Server) UpdateCommInChannel(userID int64, channelID int64, comm string) (err error) {
	commID,err := s.GetCommID(comm)
	if err != nil {
		return
	}

	set, err := s.dataBase.Prepare("UPDATE CHANNEL_USER SET comm_id = ? WHERE channel_id = ? AND user_id = ?")
	if err != nil {
		return
	}

	defer set.Close()

	_, err = set.Exec(commID, channelID, userID)
	if err != nil {
		return
	}
	return
}