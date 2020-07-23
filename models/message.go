package models

import "github.com/jinzhu/gorm"

// Message model
type Message struct {
	gorm.Model
	Body       string
	Chat       Chat `gorm:"foreignkey:ChatID"`
	ChatID     uint
	ReceiverID uint
	IsRead     bool
}
