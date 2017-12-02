package channel

type Preference struct {
	StartDate		int64		`json:"start_date"`
	Duration		int			`json:"duration"`
	Intervals		[]Interval	`json:"intervals"`
}

func (s *Server) GetPreference(userID int) (pre Preference, err error) {
	var prefID int

	pre.Intervals = make([]Interval, 0)
	get, err := s.dataBase.Prepare("SELECT preference_id ,start_date, duration_days FROM USERS AS U, PREFERENCE AS P WHERE U.user_id = ? AND P.user_id = ?")
	if err != nil {
		return
	}

	err = get.QueryRow(userID, userID).Scan(&prefID, &pre.StartDate, &pre.Duration)
	if err != nil {
		return
	}

	pre.Intervals, err = s.GetInterval(prefID)
	return
}

func (s *Server) AddPreference(p Preference, channelID int) (prefID int64, err error) {
	set, err := s.dataBase.Prepare("INSERT INTO PREFERENCE(duration_days, start_date, channel_id) VALUES(?, ?, ?)")
	if err != nil {
		return
	}

	res, err := set.Exec(p.Duration, p.StartDate, channelID)
	if err != nil {
		return
	}

	s.

	prefID, err = res.LastInsertId()
	return
}