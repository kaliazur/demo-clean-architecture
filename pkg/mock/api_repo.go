package mock

import (
	// "github.com/kaliazur/demo-clean-architecture/v2/pkg/domain"
)

type ApiRepo struct {
	apis int
}

func (a *ApiRepo) Init(config map[string]interface{}) (err error) {
	a.apis = 0
	return nil
}

func (a *ApiRepo) SetApis(nbr int) (err error) {
	a.apis = nbr
	return
}

func (a *ApiRepo) AddApis(nbr int) (err error) {
	a.apis += nbr
	return
}

func (a *ApiRepo) DelApis(nbr int) (err error) {
	a.apis -= nbr
	return
}

func (a *ApiRepo) GetApis() (nbr int) {
	return a.apis
}