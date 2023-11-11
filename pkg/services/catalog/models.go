package catalog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	pb "micromango/pkg/grpc/catalog"
)

type Manga struct {
	MangaId     uuid.UUID `json:"mangaId" gorm:"primaryKey;type:uuid"`
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	Description string    `json:"description"`
}

func (m *Manga) BeforeCreate(*gorm.DB) error {
	m.MangaId = uuid.New()
	return nil
}

func (m *Manga) ToResponse() *pb.MangaResponse {
	return &pb.MangaResponse{
		MangaId:       m.MangaId.String(),
		Title:         m.Title,
		Cover:         m.Cover,
		Description:   m.Description,
		ChapterNumber: 0,
	}
}
