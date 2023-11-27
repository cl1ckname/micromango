package catalog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	pb "micromango/pkg/grpc/catalog"
	"time"
)

type Manga struct {
	MangaId     uuid.UUID `json:"mangaId" gorm:"primaryKey;type:uuid"`
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (m *Manga) BeforeCreate(*gorm.DB) error {
	m.MangaId = uuid.New()
	m.CreatedAt = time.Now()
	return nil
}

func (m *Manga) ToResponse() *pb.MangaResponse {
	return &pb.MangaResponse{
		MangaId:     m.MangaId.String(),
		Title:       m.Title,
		Cover:       m.Cover,
		Description: m.Description,
		CreatedAt:   m.CreatedAt.String(),
	}
}
