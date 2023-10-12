package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId       uuid.UUID `json:"userId" gorm:"primaryKey"`
	PasswordHash string    `json:"passwordHash"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Picture      string    `json:"picture"`
}

func (user *User) BeforeCreate(*gorm.DB) error {
	user.UserId = uuid.New()
	return nil
}
