package Model

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	LibraryId   int    `gorm:"primary_key;autoIncrement"`
	LibraryName string `json:"LibraryName"`
	Address     string `json:"Address"`
	Phone       string `json:"Phone"`
}
