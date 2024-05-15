package Token

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	Id          int       `gorm:"primary_key;autoIncrement"`
	AccessToken string    `json:"AccessToken"`
	ExpiresAt   time.Time `json:"ExpiresAt"`
	UserID      int       `json:"User"`
	IsRevoked   bool      `json:"IsRevoked"`
}
