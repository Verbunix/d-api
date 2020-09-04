package base

type IService interface {
}

type Service struct {
	IService
	Repository IRepository
	Logger     ILogger
}

func NewBaseService(repository IRepository, logger ILogger) *Service {
	return &Service{Repository: repository, Logger: logger}
}
