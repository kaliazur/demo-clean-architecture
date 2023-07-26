package usecase

import (
	"errors"

	"github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"
)

type Interactor struct {
	UserRepo domain.UserRepo
	ApiRepo domain.ApiRepo
}

func (i *Interactor) CreateUser(
	id string,
	username string,
	firstname string,
	age int,
	admin bool,
) (createdUser *domain.User, err error) {
	return i.UserRepo.CreateUser(
		id,
		username,
		firstname,
		age,
		admin,
	)
}

func (i *Interactor) PurchaseApi(userId string, nbr int) (err error) {

	var apiNbr int = i.ApiRepo.GetApis() 

	if apiNbr <= 0{
		return errors.New("Not Enough Apis")
	}

	user, err := i.UserRepo.GetUserById(userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("User Not Found")
	}

	if user.IsAdmin() {

	// Execute the Purchase
	i.UserRepo.UserPurchaseApis(userId, nbr)
	i.ApiRepo.DelApis(nbr)
}

	return
}
