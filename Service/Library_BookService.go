package Service

import (
	"gorm-tutorial/Dao"
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type LibraryBookService struct {
	repository *Dao.LibraryBookRepository
}

func NewLibraryBookService(db *gorm.DB) *LibraryBookService {
	LibraryBookRepo := Dao.NewLibraryBookRepository(db)
	return &LibraryBookService{
		repository: LibraryBookRepo,
	}
}
func (lbs *LibraryBookService) Create(libraryBook *Model.LibraryBook) error {
	return lbs.repository.AddLibraryBook(libraryBook)
}
