package Service

import (
	"gorm-tutorial/Dao"
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type LibraryService struct {
	LibraryRepository *Dao.LibraryRepository
}

func NewLibraryService(db *gorm.DB) *LibraryService {
	lr := Dao.NewLibraryRepository(db)
	return &LibraryService{
		LibraryRepository: lr,
	}
}
func (ls *LibraryService) AddBook(library *Model.Library) error {
	return ls.LibraryRepository.AddBook(library)
}
