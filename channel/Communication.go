package channel

type Communication struct {
	Type	string	`json:"type"`
	Value	string	`json:"value"`
}

func (s *Server) GetCommunication(commID int) (com []Communication, err error) {
	com = make([]Communication, 0)
	get, err := s.dataBase.Prepare("SELECT C.val, CT.val FROM COMM AS C ,COMM_TYPE AS CT WHERE C.user_id = ? AND C.type_id = CT.type_id")
	if err != nil {
		return
	}

	rows, err := get.Query(commID)
	if rows == nil || err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		i := Communication{}
		rows.Scan(&i.Value, &i.Type)
		com = append(com, i)
	}

	return
}