package base

import "github.com/jinzhu/gorm"

type IRepository interface {
}

type Repository struct {
	IRepository
	Db     *gorm.DB
	Logger ILogger
}

func NewRepository(db *gorm.DB, logger ILogger) *Repository {
	return &Repository{Db: db, Logger: logger}
}
