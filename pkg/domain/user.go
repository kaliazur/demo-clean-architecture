package domain

type User struct {
	Id        	string
	Username  	string
	Firstname 	string
	Age       	int
	Admin   	bool
	Apis  		int
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) IsAdmin() bool {
	return u.Admin
}

func (u *User) GetApis() int {
	return u.Apis
}