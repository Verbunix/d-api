package base

type IService interface {
}

type Service struct {
	IService
	Repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{Repository: repository}
}
