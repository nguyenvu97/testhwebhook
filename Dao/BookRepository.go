package Dao

import (
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}
func (ur BookRepository) AddBook(book *Model.Book) error {
	result := ur.db.Create(book)
	return result.Error
}

func (ur *BookRepository) UpdateBook(book *Model.Book) error {
	result := ur.db.Save(book)
	return result.Error
}
func (ur *BookRepository) SreachBookCategory(bookName string) (*[]Model.Book, error) {
	var book []Model.Book
	result := ur.db.Where("BookCategory = ?", bookName).First(&book)
	return &book, result.Error
}
func (ur *BookRepository) SreachBook(bookName string) (*[]Model.Book, error) {
	var book []Model.Book
	result := ur.db.Where("BookCategory = ?", bookName).First(&book)
	return &book, result.Error
}

func (ur *BookRepository) GetAllBook() (*[]Model.Book, error) {
	var book []Model.Book
	result := ur.db.Find(&book)
	return &book, result.Error
}
func (ur *BookRepository) FindByBookId(id int) (*Model.Book, error) {
	var book Model.Book
	result := ur.db.Find(id, &book)
	return &book, result.Error
}
