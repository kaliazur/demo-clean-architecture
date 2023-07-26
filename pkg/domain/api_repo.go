package domain

type ApiRepo interface {
	Init(config map[string]interface{}) (err error)
	SetApis(nbr int) (err error)
	AddApis(nbr int) (err error)
	DelApis(nbr int) (err error)
	GetApis() (nbr int)
}
