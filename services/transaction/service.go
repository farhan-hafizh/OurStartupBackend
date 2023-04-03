package transaction

type Service interface{}

type service struct {
	repository Repository
}

func CreateService(repository Repository) *service {
	return &service{repository}
}
