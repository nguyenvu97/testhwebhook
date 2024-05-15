package Dao

import (
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type LibraryRepository struct {
	db *gorm.DB
}

func NewLibraryRepository(db *gorm.DB) *LibraryRepository {
	return &LibraryRepository{db: db}
}
func (lr LibraryRepository) AddBook(libraryRepository *Model.Library) error {
	result := lr.db.Create(libraryRepository)
	return result.Error
}
