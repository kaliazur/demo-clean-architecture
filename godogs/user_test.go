package features

import (
	"fmt"
	"github.com/cucumber/godog"

	// domain "github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"
	mock "github.com/kaliazur/demo-clean-architecture/v2/pkg/mock"
	usecase "github.com/kaliazur/demo-clean-architecture/v2/pkg/usecase"
)

var interactor = &usecase.Interactor{
	UserRepo: &mock.UserRepo{},
	ApiRepo: &mock.ApiRepo{},
}

func apis(arg1 int) error {
	return interactor.ApiRepo.SetApis(arg1)
}

func nonAdminUserWithIDUsernameFirstnameAge(arg1, arg2, arg3 string, arg4 int) error {
	_, err := interactor.CreateUser(
		arg1,
		arg2,
		arg3,
		arg4,
		false,
	)
	return err
}

func adminUserWithIDUsernameFirstnameAge(arg1, arg2, arg3 string, arg4 int) error {
	_, err := interactor.CreateUser(
		arg1,
		arg2,
		arg3,
		arg4,
		true,
	)
	return err
}

func iAddApis(arg1 int) error {
	return interactor.ApiRepo.AddApis(arg1)
}

func userWithIDPurchaseApi(arg1 string, arg2 int) error {
	return interactor.PurchaseApi(arg1, arg2)
}

func thereShouldBeRemainingApi(arg1 int) error {
	if arg1 == interactor.ApiRepo.GetApis() {
		return nil
	}
	return fmt.Errorf("expected %d api to be remaining, but there is %d", arg1, interactor.ApiRepo.GetApis())
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^non admin user with ID "([^"]*)", username "([^"]*)", firstname "([^"]*)", age (\d+)$`, nonAdminUserWithIDUsernameFirstnameAge)
	ctx.Step(`^user with ID "([^"]*)" purchase (\d+) api$`, userWithIDPurchaseApi)
	ctx.Step(`^(\d+) apis$`, apis)
	ctx.Step(`^I add (\d+) apis$`, iAddApis)
	ctx.Step(`^there should be (\d+) remaining api$`, thereShouldBeRemainingApi)
	ctx.Step(`^admin user with ID "([^"]*)", username "([^"]*)", firstname "([^"]*)", age (\d+)$`, adminUserWithIDUsernameFirstnameAge)
}
