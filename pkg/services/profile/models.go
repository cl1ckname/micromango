package profile

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	pb "micromango/pkg/grpc/profile"
	"time"
)

type Profile struct {
	UserId    uuid.UUID `json:"userId" gorm:"primaryKey;type:uuid"`
	Username  string    `json:"username"`
	Picture   string    `json:"picture"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
}

func (m *Profile) BeforeCreate(*gorm.DB) error {
	m.UserId = uuid.New()
	m.CreatedAt = time.Now()
	return nil
}

func (m *Profile) ToResponse() *pb.ProfileResponse {
	return &pb.ProfileResponse{
		UserId:    m.UserId.String(),
		Username:  m.Username,
		Picture:   m.Picture,
		Bio:       m.Bio,
		CreatedAt: m.CreatedAt.String(),
	}
}
