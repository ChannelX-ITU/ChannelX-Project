package channel

type Restriction struct {
	prefID		int
	TypeID		string	`json:"comm_type, omitempty"`
	Val			string 	`json:"value"`
	ContOp		string	`json:"operator"`
	ContType	string	`json:"type"`
}

func (s *Server) AddRestriction(in []Restriction, prefID int64) (err error) {
	set, err := s.dataBase.Prepare("INSERT INTO RESTRICTION(preference_id, type_id, val, cont_op, cont_type) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return
	}

	var t int
	for _, val := range in {
		if val.TypeID == "SMS" {
			t = 2
		} else {
			t = 1
		}
		_, err = set.Exec(prefID, t, val.Val, val.ContOp, val.ContType)
	}

	return
}

func (s *Server) GetRestrictions(prefID int64) (inf []Restriction, err error) {
	inf = make([]Restriction, 0)
	get, err := s.dataBase.Prepare("SELECT CT.val, R.val, R.cont_op, R.cont_type FROM RESTRICTION AS R, COMM_TYPE AS CT WHERE R.preference_id = ? AND CT.type_id = R.type_id")
	if err != nil {
		return
	}

	rows, err := get.Query(prefID)
	if rows == nil || err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		i := Restriction{}
		rows.Scan(&i.TypeID, &i.Val, &i.ContOp, &i.ContType)
		inf = append(inf, i)
	}

	return
}