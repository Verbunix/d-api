package base

import "gorm.io/gorm"

type IRepository interface {
}

type Repository struct {
	IRepository
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}
