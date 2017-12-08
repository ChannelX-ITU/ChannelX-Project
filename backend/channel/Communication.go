package channel

type Communication struct {
	Type	string	`json:"comm_type"`
	Value	string	`json:"value"`
}

func (s *Server) GetCommunication(commID int64) (com []Communication, err error) {
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

func (s *Server) GetCommType(commType string) (id uint64, err error) {
	get, err := s.dataBase.Prepare("SELECT type_id FROM COMM_TYPE WHERE val = ?")
	if err != nil {
		return
	}

	err = get.QueryRow(commType).Scan(&id)
	if err != nil {
		return
	}

	defer get.Close()

	return
}

func (s *Server) SendMessage(mes SendMessage, comm Communication) {
	switch comm.Type {
	case "SMS":
		return

	case "EMAIL":
		s.mailMan.Send(Message{To:comm.Value, Sub:mes.Subject, Msg:mes.Message})
	}
}