package Model

import "gorm.io/gorm"

type LibraryBook struct {
	gorm.Model
	Id        int `gorm:"primary_key;autoIncrement"`
	LibraryId int `json:"LibraryId"`
	Book      int `json:"Book"`
	Quantity  int `json:"Quantity"`
}
