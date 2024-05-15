package Model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	id           int     `gorm:"primary_key;autoIncrement"`
	Title        string  `json:"Title"`
	Author       string  `json:"Author"`
	BookName     string  `json:"BookName"`
	BookPrice    float64 `json:"BookPrice"`
	BookCategory string  `json:"BookCategory"`
	QuantityBook int     `json:"QuantityBook"`
}
