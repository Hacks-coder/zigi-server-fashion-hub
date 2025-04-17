package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"not null" json:"username"`
	Email        string `gorm:"not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"passowrd,omitempty"`
}
