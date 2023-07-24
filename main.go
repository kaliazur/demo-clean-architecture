package main

import (
	"fmt"

	domain "github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"

	mock "github.com/kaliazur/demo-clean-architecture/v2/pkg/mock"
	// infrastructure "github.com/kaliazur/demo-clean-architecture/v2/pkg/infrastructure"

	usecase "github.com/kaliazur/demo-clean-architecture/v2/pkg/usacase"
)

var interactor = &usecase.Interactor{
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

func main() {

	fmt.Println("main()")

	user, err := interactor.CreateUser(domain.User{
		Username:  "kaliazur",
		Firstname: "Pierre",
		Age:       27,
		IsAdmin:   true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("User is: %+v\n", user)

	user, err = interactor.UserRepo.GetUserById(user.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User is: %+v\n", user)

}
