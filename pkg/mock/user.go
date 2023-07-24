package mock

import (
	"errors"
	"fmt"

	"github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"

	"github.com/google/uuid"
)

type UserRepo struct {
}

var usersMap map[string]*domain.User = map[string]*domain.User{}

func (u *UserRepo) Init(config map[string]interface{}) (err error) {

	fmt.Println("Init()")

	return nil
}

func (u *UserRepo) CreateUser(user domain.User) (createdUser *domain.User, err error) {
	fmt.Println("CreateUser()")

	if len(user.Username) == 0 {

		err = errors.New("field 'username' required")

	} else if len(user.Firstname) == 0 {

		err = errors.New("field 'firstname' required")

	} else if user.Age <= 0 {

		err = errors.New("field 'Age' required")

	} else {

		user.Id = uuid.NewString()
		usersMap[user.Id] = &user
		createdUser = &user
	}

	return
}

func (u *UserRepo) GetUserById(id string) (user *domain.User, err error) {
	fmt.Printf("GetUserById(id: %s)\n", id)

	if _, ok := usersMap[id]; ok {
		user = usersMap[id]
	}

	return
}
