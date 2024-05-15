package Model

import "gorm.io/gorm"

type UserBook struct {
	gorm.Model
	UserBookId   int     `gorm:"primary_Key;autoIncrement"`
	UserId       int     `json:"UserId"`
	BookId       int     `json:"BookId"`
	QuantityBook int     `json:"QuantityBook"`
	MoneyBook    float64 `json:"MoneyBook"`
}
