package domain

type UserRepo interface {
	Init(config map[string]interface{}) (err error)
	CreateUser(
		id string,
		username string,
		firstname string,
		age int,
		isAdmin bool,
	) (createdUser *User, err error)
	GetUserById(id string) (user *User, err error)
	UserPurchaseApis(id string, nbr int) (err error)
}
