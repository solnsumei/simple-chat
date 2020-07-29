package utils

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// UserInput request
type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate users
func (input UserInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Email, validation.Required, is.EmailFormat),
		validation.Field(&input.Password, validation.Required, validation.Length(5, 50)),
	)
}

// MessageInput request
type MessageInput struct {
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Message    string `json:"message"`
	ChatID string `json:"chatId"`
}

// Validate messages
func (msgInput MessageInput) Validate() error {
	return validation.ValidateStruct(&msgInput,
		validation.Field(&msgInput.SenderID, validation.Required),
		validation.Field(&msgInput.ReceiverID, validation.Required),
		validation.Field(&msgInput.Message, validation.Required),
		validation.Field(&msgInput.Message, validation.Required),
	)
}
