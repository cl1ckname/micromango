package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	pb "micromango/pkg/grpc/user"
)

type User struct {
	UserId       uuid.UUID `json:"userId" gorm:"primaryKey"`
	PasswordHash string    `json:"passwordHash"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Picture      string    `json:"picture"`
}

func (u *User) BeforeCreate(*gorm.DB) error {
	u.UserId = uuid.New()
	return nil
}

func (u User) ToPb() *pb.UserResponse {
	return &pb.UserResponse{
		UserId:   u.UserId.String(),
		Username: u.Username,
		Email:    u.Email,
		Picture:  u.Picture,
	}
}
