package main

import (
	"fmt"

	// domain "github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"

	mock "github.com/kaliazur/demo-clean-architecture/v2/pkg/mock"
	// infrastructure "github.com/kaliazur/demo-clean-architecture/v2/pkg/infrastructure"

	usecase "github.com/kaliazur/demo-clean-architecture/v2/pkg/usecase"
)

var interactor = &usecase.Interactor{
	UserRepo: &mock.UserRepo{},
	ApiRepo: &mock.ApiRepo{},

	// UserRepo: &infrastructure.UserRepo{},
}

func init() {

	/*
	** Init User Repo
	*/
	err := interactor.UserRepo.Init(map[string]interface{}{
		"key": "value",

		// "AWS_ACCESS_KEY_ID":     `AKIAKEJDDKJFDANNWPDRJMI`,
		// "AWS_SECRET_ACCESS_KEY": `Kvzcv1Tc8Knl9yw37DKcdkfjdruecjZpVyfV4xaSX`,
	})
	if err != nil {
		panic(err)
	}

}

func main() {

	fmt.Println("main()")

	interactor.ApiRepo.SetApis(5)

	user, err := interactor.CreateUser(
		"42",
		"kaliazur",
		"Pierre",
		30,
		true,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created User is: %+v\n", user)

	interactor.PurchaseApi(user.GetId(), 4)

	fmt.Printf("User %s apis: %d\n", user.GetId(), user.GetApis())
	fmt.Printf("Apis remaining in apis store: %d\n", interactor.ApiRepo.GetApis())



}
