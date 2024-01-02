package catalog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/catalog"
	"time"
)

type Manga struct {
	MangaId     uuid.UUID `json:"mangaId" gorm:"primaryKey;type:uuid"`
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	Description string    `json:"description"`
	Genres      []Genre   `json:"genres" gorm:"many2many:manga_genres"`
	Rate        float32   `json:"rate"`
	Rates       uint64    `json:"rates"`
	Likes       uint64    `json:"likes"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedBy   uuid.UUID `json:"createdBy"`
	UpdatedBy   uuid.UUID `json:"updatedBy"`
}

func (m *Manga) BeforeCreate(*gorm.DB) error {
	m.MangaId = uuid.New()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Manga) BeforeSave(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Manga) ToResponse() *pb.MangaResponse {
	return &pb.MangaResponse{
		MangaId:     m.MangaId.String(),
		Title:       m.Title,
		Cover:       m.Cover,
		Description: m.Description,
		CreatedAt:   m.CreatedAt.String(),
		Genres:      utils.Map(m.Genres, func(g Genre) uint32 { return uint32(g.GenreId) }),
		Rate:        m.Rate,
		Rates:       m.Rates,
		Likes:       m.Likes,
	}
}

type Genre struct {
	GenreId int    `json:"genreId" gorm:"primaryKey;autoIncrement"`
	Label   string `json:"label"`
}
