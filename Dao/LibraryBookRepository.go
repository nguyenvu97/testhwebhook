package Dao

import (
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type LibraryBookRepository struct {
	db *gorm.DB
}

func NewLibraryBookRepository(db *gorm.DB) *LibraryBookRepository {
	return &LibraryBookRepository{db: db}
}
func (lbr *LibraryBookRepository) AddLibraryBook(libraryBook *Model.LibraryBook) error {
	result := lbr.db.Create(libraryBook)
	return result.Error
}
