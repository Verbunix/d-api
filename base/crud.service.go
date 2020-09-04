package base

type CrudServiceInterface interface {
	IService
	GetModel() IEntity
	FindOne(id uint) (IEntity, error)
	FindAll() ([]IEntity, error)
	Create(item IEntity) IEntity
	Update(item IEntity) IEntity
	Delete(id uint) error
}

type CrudService struct {
	*Service
	Repository ICrudRepository
}

func NewCrudService(repository ICrudRepository, logger ILogger) *CrudService {
	service := NewBaseService(repository, logger)
	return &CrudService{service, repository}
}

func (c CrudService) GetModel() IEntity {
	return c.Repository.GetModel()
}

func (c CrudService) FindOne(id uint) (IEntity, error) {
	return c.Repository.FindOne(id)
}

func (c CrudService) FindAll() ([]IEntity, error) {
	return c.Repository.FindAll()
}

func (c CrudService) Create(item IEntity) IEntity {
	return c.Repository.Create(item)
}

func (c CrudService) Update(item IEntity) IEntity {
	return c.Repository.Update(item)
}

func (c CrudService) Delete(id uint) error {
	return c.Repository.Delete(id)
}
