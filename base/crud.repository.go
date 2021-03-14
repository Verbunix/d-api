package base

import (
	"reflect"

	"gorm.io/gorm"
)

type ICrudRepository interface {
	IRepository
	GetModel() IEntity
	FindOne(id uint) (IEntity, error)
	FindAll() ([]IEntity, error)
	Create(item IEntity) IEntity
	Update(item IEntity) IEntity
	Delete(id uint) error
}

type CrudRepository struct {
	ICrudRepository
	*Repository
	Model IEntity // Dynamic typing
}

func NewCrudRepository(db *gorm.DB, model IEntity) *CrudRepository {
	repo := NewRepository(db)
	return &CrudRepository{
		Repository: repo,
		Model:      model,
	}
}

func (c CrudRepository) GetModel() IEntity {
	return c.Model
}

func (c CrudRepository) FindOne(id uint) (IEntity, error) {
	item := reflect.New(reflect.TypeOf(c.GetModel()).Elem()).Interface()
	err := c.Db.First(item, id).Error
	return item, err
}

func (c CrudRepository) FindAll() (IEntity, error) {
	item := reflect.New(reflect.TypeOf(c.GetModel()).Elem()).Interface()
	err := c.Db.Find(item).Error
	return item, err
}

func (c CrudRepository) Create(item IEntity) IEntity {
	c.Db.Create(item)
	return item
}

func (c CrudRepository) Update(item IEntity) IEntity {
	c.Db.Save(item)
	return item
}

func (c CrudRepository) Delete(id uint) error {
	item, err := c.FindOne(id)
	if err != nil {
		return err
	}
	c.Db.Delete(item)
	return nil
}
