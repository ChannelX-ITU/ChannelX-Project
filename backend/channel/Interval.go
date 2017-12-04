package channel

type Interval struct {
	Start		int		`json:"start"`
	Duration	int		`json:"duration"`
}

func (s *Server) GetInterval(prefID int) (inf []Interval, err error) {
	inf = make([]Interval, 0)
	get, err := s.dataBase.Prepare("SELECT start_time_in_minutes, duration FROM INTER WHERE preference_id = ?")
	if err != nil {
		return
	}

	rows, err := get.Query(prefID)
	if rows == nil || err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		i := Interval{}
		rows.Scan(&i.Start, &i.Duration)
		inf = append(inf, i)
	}

	return
}

func (s *Server) AddInterval(in []Interval, prefID int64) (err error) {
	set, err := s.dataBase.Prepare("INSERT INTO INTER(preference_id, start_time_in_minutes, duration) VALUES(?, ?, ?)")
	if err != nil {
		return
	}

	for _, val := range in {
		_, err = set.Exec(prefID, val.Start, val.Duration)
	}

	return
}