package mock

import (
	"github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"
)

type UserRepo struct {
}

var usersMap map[string]*domain.User = map[string]*domain.User{}

func (u *UserRepo) Init(config map[string]interface{}) (err error) {

	// fmt.Println("Init()")

	return nil
}

func (u *UserRepo) CreateUser(
	id string,
	username string,
	firstname string,
	age int,
	admin bool,
) (createdUser *domain.User, err error) {
	usersMap[id] = &domain.User{
		Id: id,
		Username: username,
		Firstname: firstname,
		Age: age,
		Admin: admin, 
	}
	createdUser = usersMap[id]
	return 
}

func (u *UserRepo) GetUserById(id string) (user *domain.User, err error) {
	// fmt.Printf("GetUserById(id: %s)\n", id)

	if _, ok := usersMap[id]; ok {
		user = usersMap[id]
	}

	return
}

func (u *UserRepo) UserPurchaseApis(id string, nbr int) (err error) {
	if _, ok := usersMap[id]; ok {
		usersMap[id].Apis += nbr
	}

	return
}
