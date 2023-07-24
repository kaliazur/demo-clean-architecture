package usecase

import (
	"fmt"
	"testing"

	domain "github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"
	mock "github.com/kaliazur/demo-clean-architecture/v2/pkg/mock"
	// infrastructure "github.com/kaliazur/demo-clean-architecture/v2/pkg/infrastructure"
)

var interactor = &Interactor{
	UserRepo: &mock.UserRepo{},
	// UserRepo: &infrastructure.UserRepo{},
}

func init() {

	/*
	** Init User Repo
	 */
	err := interactor.UserRepo.Init(map[string]interface{}{
		"key": "value",

		// "AWS_ACCESS_KEY_ID":     `AKIAS4VMHDANNWPDRJMI`,
		// "AWS_SECRET_ACCESS_KEY": `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`,
	})
	if err != nil {
		panic(err)
	}

}

func TestCreateUser(t *testing.T) {

	t.Run("purchase with not admin user", func(t *testing.T) {

		// We create the user
		user, err := interactor.UserRepo.CreateUser(domain.User{
			Username:  "kaliazur",
			Firstname: "Pierre",
			Age:       27,
			IsAdmin:   false,
		})
		if err != nil {
			t.Error("Error creating the user", err)
		}
		fmt.Printf("User is: %+v\n", user)

		user, err = interactor.UserRepo.GetUserById(user.Id)
		if err != nil {
			t.Error("Error getting the user", err)
		}
		if user == nil {
			t.Error("Previously created user not found")
		}
		fmt.Printf("User is: %+v\n", user)

		err = interactor.PurchaseApi(user.Id, "api_id_test")
		if err == nil || err.Error() != "Unauthorized" {
			t.Fail()
		}

	})

	t.Run("purchase with admin user", func(t *testing.T) {

		// We create the user
		user, err := interactor.UserRepo.CreateUser(domain.User{
			Username:  "kaliazur",
			Firstname: "Pierre",
			Age:       27,
			IsAdmin:   true,
		})
		if err != nil {
			t.Error("Error creating the user", err)
		}
		fmt.Printf("User is: %+v\n", user)

		user, err = interactor.UserRepo.GetUserById(user.Id)
		if err != nil {
			t.Error("Error getting the user", err)
		}
		if user == nil {
			t.Error("Previously created user not found")
		}
		fmt.Printf("User is: %+v\n", user)

		err = interactor.PurchaseApi(user.Id, "api_id_test")
		if err != nil {
			t.Fail()
		}

	})

}
