package channel

import (
	"net/url"
	"strings"
	"net/http"
)

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

	defer get.Close()

	err = get.QueryRow(commType).Scan(&id)
	if err != nil {
		return
	}

	return
}

func (s *Server) SendMessage(mes SendMessage, comm Communication) {
	switch comm.Type {
	case "SMS":
		SendSMS(mes, comm)
		return

	case "EMAIL":
		s.mailMan.Send(Message{To:comm.Value, Sub:mes.Subject, Msg:mes.Message})
		return
	}
}

func SendSMS(mes SendMessage, comm Communication) {
	accountSid := "AC5122131b79a439f9fe9d00b12e731c9f"
	authToken := "e10f95591a3091b840e046137fa4f495"
	from := "+12109636969"
	to := comm.Value
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/AC5122131b79a439f9fe9d00b12e731c9f/Messages.json"
	msgData := url.Values{}
	msgData.Set("To",to)
	msgData.Set("From",from)
	msgData.Set("Body",mes.Subject + " - " + mes.Message)
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client.Do(req)
}

func (s *Server) GetCommTypeName(commType string) (id string, err error) {
	get, err := s.dataBase.Prepare("SELECT CT.val FROM COMM_TYPE AS CT, COMM AS C WHERE C.val = ? AND CT.type_id = C.type_id")
	if err != nil {
		return
	}

	defer get.Close()

	err = get.QueryRow(commType).Scan(&id)
	if err != nil {
		return
	}

	return
}