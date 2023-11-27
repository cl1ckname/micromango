package reading

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MangaContent struct {
	MangaId  uuid.UUID `json:"mangaId" gorm:"primaryKey;type:uuid"`
	Chapters []Chapter `json:"chapters"  gorm:"foreignKey:MangaId"`
}

func (content *MangaContent) BeforeCreate(*gorm.DB) error {
	if content.MangaId == uuid.Nil {
		content.MangaId = uuid.New()
	}
	return nil
}

type Chapter struct {
	ChapterId uuid.UUID `json:"chapterId" gorm:"primaryKey;type:uuid"`
	MangaId   uuid.UUID `json:"mangaId"`
	Title     string    `json:"title"`
	Number    float32   `json:"number"`
	Pages     []Page    `json:"pages"`
}

func (chapter *Chapter) BeforeCreate(*gorm.DB) error {
	if chapter.ChapterId == uuid.Nil {
		chapter.ChapterId = uuid.New()
	}
	return nil
}

type Page struct {
	PageId    uuid.UUID `json:"pageId" gorm:"primaryKey;type:uuid"`
	ChapterId uuid.UUID `json:"chapterId"`
	Number    uint32    `json:"number"`
	Image     string    `json:"image"`
}

func (page *Page) BeforeCreate(*gorm.DB) error {
	if page.ChapterId == uuid.Nil {
		page.ChapterId = uuid.New()
	}
	return nil
}
