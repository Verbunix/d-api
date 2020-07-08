package services

import (
	"dating-api/models"
	"github.com/jinzhu/gorm"
	"reflect"
)

type IService interface{}

type Service struct {
	IService
	db *gorm.DB
}

func NewService(db *gorm.DB) IService {
	return &Service{db: db}
}

type IParameters interface{}

type PaginationParameters struct {
	Page     int
	PageSize int
}

type CrudParameters struct {
	IParameters
	*PaginationParameters
}

type ICrud interface {
	IService
	GetModel() models.IModel
	Find(id uint) (models.IModel, error)
	List(parameters IParameters) (models.IModel, error)
	Create(item models.IModel) models.IModel
	Update(item models.IModel) models.IModel
	Delete(id uint) error
}

type IQueryBuilder interface {
	ListQuery(parameters IParameters) *gorm.DB
}

type QueryBuilder struct {
	db *gorm.DB
	IQueryBuilder
}

type Crud struct {
	ICrud
	*Service
	model        models.IModel // Dynamic typing
	QueryBuilder IQueryBuilder
}

func NewCrudService(db *gorm.DB, model models.IModel, queryBuilder IQueryBuilder) ICrud {
	service := NewService(db).(*Service)
	return &Crud{
		Service:      service,
		model:        model,
		QueryBuilder: queryBuilder,
	}
}

func (c QueryBuilder) paginationQuery(parameters IParameters) *gorm.DB {
	query := c.db

	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		return query
	}

	var page int64
	page = 0
	pageValue := val.FieldByName("Page")
	if pageValue.IsValid() || pageValue.Kind() == reflect.Int {
		page = pageValue.Int()
	}

	var pageSize int64
	pageSize = 20 // DefaultPageSize
	pageSizeValue := val.FieldByName("PageSize")
	if pageSizeValue.IsValid() || pageSizeValue.Kind() == reflect.Int {
		pageSize = pageSizeValue.Int()
	}

	limit := pageSize
	offset := page * pageSize
	query = query.Offset(offset).Limit(limit)

	return query
}

func (c QueryBuilder) ListQuery(parameters IParameters) *gorm.DB {
	return c.paginationQuery(parameters)
}

func (c Crud) GetModel() models.IModel {
	return c.model
}

func (c Crud) Find(id uint) (models.IModel, error) {
	item := reflect.New(reflect.TypeOf(c.GetModel()).Elem()).Interface()
	err := c.db.First(item, id).Error
	return item, err
}

func (c Crud) List(parameters IParameters) (models.IModel, error) {
	items := reflect.New(reflect.SliceOf(reflect.TypeOf(c.GetModel()).Elem())).Interface()
	query := c.QueryBuilder.ListQuery(parameters)
	err := query.Find(items).Error
	return items, err
}

func (c Crud) Create(item models.IModel) models.IModel {
	c.db.Create(item)
	return item
}

func (c Crud) Update(item models.IModel) models.IModel {
	c.db.Save(item)
	return item
}

func (c Crud) Delete(id uint) error {
	item, err := c.Find(id)
	if err != nil {
		return err
	}
	c.db.Delete(item)
	return nil
}
