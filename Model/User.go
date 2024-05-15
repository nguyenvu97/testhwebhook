package Model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id        int    `gorm:"primary_key;autoIncrement"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}
type UserDto struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

func MapUserToDTO(user *User) *UserDto {
	return &UserDto{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

}
