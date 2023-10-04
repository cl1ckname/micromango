package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Manga struct {
	MangaId     uuid.UUID `json:"mangaId" gorm:"primaryKey;type:uuid"`
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	Description string    `json:"description"`
}

func (manga *Manga) BeforeCreate(*gorm.DB) error {
	manga.MangaId = uuid.New()
	return nil
}
