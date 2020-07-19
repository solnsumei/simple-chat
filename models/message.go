package models

import "github.com/jinzhu/gorm"

// Message model
type Message struct {
	gorm.Model
	Body       string
	Sender     User `gorm:"foreignkey:SenderID"`
	SenderID   uint
	Receiver   User `gorm:"foreignkey:ReceiverID"`
	ReceiverID uint
}
