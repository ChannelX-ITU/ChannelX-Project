package channel

type Restriction struct {
	prefID		int
	typeID		int
	Val			string 	`json:"value"`
	ContOp		string	`json:"operator"`
	ContType	string	`json:"type"`
}
