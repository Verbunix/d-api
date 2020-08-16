package base

import (
	"github.com/jinzhu/gorm"
)

type InterfaceEntity interface {
}

// LOGGER
type LoggerInterface interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// BASE REPOSITORY
type BaseRepositoryInterface interface {
}

type BaseRepository struct {
	BaseRepositoryInterface
	Db     *gorm.DB
	Logger LoggerInterface
}

func NewBaseRepository(db *gorm.DB, logger LoggerInterface) *BaseRepository {
	return &BaseRepository{Db: db, Logger: logger}
}

// BASE SERVICE
type BaseServiceInterface interface {
}

type BaseService struct {
	BaseServiceInterface
	Repository BaseRepositoryInterface
	Logger     LoggerInterface
}

func NewBaseService(repository BaseRepositoryInterface, logger LoggerInterface) *BaseService {
	return &BaseService{Repository: repository, Logger: logger}
}

// CRUD SERVICE
type CrudServiceInterface interface {
	BaseServiceInterface
	GetModel() InterfaceEntity
	GetItem(id uint) (InterfaceEntity, error)
	GetList(parameters ListParametersInterface) ([]InterfaceEntity, error)
	Create(item InterfaceEntity) InterfaceEntity
	Update(item InterfaceEntity) InterfaceEntity
	Delete(id uint) error
}

type CrudService struct {
	*BaseService
	Repository CrudRepositoryInterface
}

func NewCrudService(repository CrudRepositoryInterface, logger LoggerInterface) *CrudService {
	service := NewBaseService(repository, logger)
	return &CrudService{service, repository}
}

func (c CrudService) GetModel() InterfaceEntity {
	return c.Repository.GetModel()
}

func (c CrudService) GetItem(id uint) (InterfaceEntity, error) {
	return c.Repository.Find(id)
}

func (c CrudService) GetList(parameters ListParametersInterface) ([]InterfaceEntity, error) {
	return c.Repository.List(parameters)
}

func (c CrudService) Create(item InterfaceEntity) InterfaceEntity {
	return c.Repository.Create(item)
}

func (c CrudService) Update(item InterfaceEntity) InterfaceEntity {
	return c.Repository.Update(item)
}

func (c CrudService) Delete(id uint) error {
	return c.Repository.Delete(id)
}
