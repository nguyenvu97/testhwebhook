package Service

import (
	"errors"
	"gorm-tutorial/Dao"
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type UserBookService struct {
	userBookRepo *Dao.UserBookReposiroty
	bookRepo     *Dao.BookRepository
	userRepo     *Dao.UserRepository
}

func NewUserBookService(db *gorm.DB) *UserBookService {
	ubr := Dao.NewUserBookReposiroty(db)
	bookRepo := Dao.NewBookRepository(db)
	userRepo := Dao.NewUserRepository(db)
	return &UserBookService{
		userBookRepo: ubr,
		bookRepo:     bookRepo,
		userRepo:     userRepo,
	}
}
func (ub *UserBookService) CreateUserBook(email string, bookIds []int, QuantityBook int) (*Model.UserBook, error) {
	for _, bookId := range bookIds {
		result, err := ub.bookRepo.FindByBookId(bookId)
		if err != nil {
			return nil, errors.New("User not found")
		}

		user, err := ub.userRepo.FindByEmail(email)
		if err != nil {
			return nil, errors.New("User not found")
		}
		var userBook Model.UserBook
		userBook.UserBookId = bookId
		userBook.QuantityBook = QuantityBook
		userBook.UserBookId = int(user.ID)
		userBook.MoneyBook = (float64(QuantityBook) * result.BookPrice) * 10 / 90

		return &userBook, nil
	}
	return nil, errors.New("No books found")

}
