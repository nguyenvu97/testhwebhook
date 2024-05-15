package Dao

import (
	"golang.org/x/crypto/bcrypt"
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) AddUser(user *Model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	result := ur.db.Create(user)
	return result.Error
}
func (ur *UserRepository) FindByID(id uint) (*Model.User, error) {
	var user Model.User
	result := ur.db.First(&user, id)
	return &user, result.Error
}
func (ur *UserRepository) Update(user *Model.User) error {
	result := ur.db.Save(user)
	return result.Error
}
func (ur *UserRepository) SreachUser(email string) (*Model.User, error) {
	var user Model.User
	result := ur.db.Where("Email = ?", email).First(&user)
	return &user, result.Error
}
func (ur *UserRepository) GetAllUser() (*[]Model.User, error) {
	var user []Model.User
	result := ur.db.Find(&user)
	return &user, result.Error
}
func (ur *UserRepository) FindByEmail(email string) (*Model.User, error) {
	var user Model.User
	result := ur.db.Where("Email = ?", email).First(&user)
	return &user, result.Error
}
