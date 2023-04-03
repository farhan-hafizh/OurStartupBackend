package transaction

import "gorm.io/gorm"

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) *repository {
	return &repository{db}
}
