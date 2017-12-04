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