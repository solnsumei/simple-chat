package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"-"`
}
