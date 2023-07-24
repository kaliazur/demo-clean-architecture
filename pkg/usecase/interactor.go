package usecase

import (
	"errors"
	"fmt"

	"github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"
)

type Interactor struct {
	UserRepo domain.UserRepo
}

func (i *Interactor) CreateUser(user domain.User) (createdUser *domain.User, err error) {
	return i.UserRepo.CreateUser(user)
}

func (i *Interactor) PurchaseApi(userId string, apiId string) error {

	user, err := i.UserRepo.GetUserById(userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New(fmt.Sprintf("User Not Found: %s", userId))
	}

	if user.IsAdmin == true {
		// Execute the Purchase

		fmt.Printf("Purchase done by user_id: %s api_id: %s\n", userId, apiId)

		return nil
	} else {

		return errors.New("Unauthorized")

	}

}
