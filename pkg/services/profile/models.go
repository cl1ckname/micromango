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

func (m *Profile) ToResponse() *pb.Response {
	return &pb.Response{
		UserId:    m.UserId.String(),
		Username:  m.Username,
		Picture:   m.Picture,
		Bio:       m.Bio,
		CreatedAt: m.CreatedAt.String(),
	}
}

type ListRecord struct {
	UserId    uuid.UUID   `json:"userId" gorm:"primaryKey;type:uuid"`
	MangaId   uuid.UUID   `json:"mangaId" gorm:"primaryKey;type:uuid"`
	ListName  pb.ListName `json:"listName" gorm:"type:uint"`
	CreatedAt time.Time   `json:"createdAt"`
}

func (lr *ListRecord) BeforeCreate(*gorm.DB) error {
	lr.CreatedAt = time.Now()
	return nil
}
