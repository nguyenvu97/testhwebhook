package Dao

import (
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type UserBookReposiroty struct {
	db *gorm.DB
}

func NewUserBookReposiroty(db *gorm.DB) *UserBookReposiroty {
	return &UserBookReposiroty{db: db}
}
func (ub *UserBookReposiroty) AddUserBook(userbook *Model.UserBook) error {
	result := ub.db.Create(userbook)
	return result.Error
}
