package channel

type User struct {
	eMail		string
	userId		int
	PrefId		int
	hashPass	string
}

func (u *User) SetMail(eMail string) {
	u.eMail = eMail
}

func (u *User) SetHashPass( hashPass string) {
	u.hashPass = hashPass
}

func (u *User) SetPrefId ( prefId int ){
	u.PrefId = prefId
}