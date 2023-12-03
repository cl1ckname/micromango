package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId       uuid.UUID `json:"userId" gorm:"primaryKey"`
	PasswordHash string    `json:"passwordHash"`
	Email        string    `json:"email"`
}

func (u *User) BeforeCreate(*gorm.DB) error {
	u.UserId = uuid.New()
	return nil
}
