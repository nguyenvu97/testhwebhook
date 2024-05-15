package Service

import (
	"errors"
	"gorm-tutorial/Dao"
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type BookService struct {
	bookRepository *Dao.BookRepository
}

func NewBookService(db *gorm.DB) *BookService {
	Repository := Dao.NewBookRepository(db)
	return &BookService{
		bookRepository: Repository,
	}
}
func (bs *BookService) AddBook(book *Model.Book) error {
	return bs.bookRepository.AddBook(book)
}
func (bs *BookService) UpdateBook(book *Model.Book) error {
	return bs.bookRepository.UpdateBook(book)
}

func (bs *BookService) SreachBook(bookName string) (*[]Model.Book, error) {
	resultByCategory, err := bs.bookRepository.SreachBookCategory(bookName)
	if err != nil {
		return nil, err
	}
	if resultByCategory != nil {
		return resultByCategory, nil
	}
	resultByName, err := bs.bookRepository.SreachBook(bookName)
	if err != nil {
		return nil, err
	}
	if resultByName == nil {
		return nil, errors.New("can not find by bookName")
	}
	return resultByName, nil

}
