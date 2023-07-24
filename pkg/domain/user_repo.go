package domain

type UserRepo interface {
	Init(config map[string]interface{}) (err error)

	CreateUser(user User) (createdUser *User, err error)

	GetUserById(id string) (user *User, err error)
}
