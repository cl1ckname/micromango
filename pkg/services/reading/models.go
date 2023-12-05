package reading

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Chapter struct {
	ChapterId uuid.UUID `json:"chapterId" gorm:"primaryKey;type:uuid"`
	MangaId   uuid.UUID `json:"mangaId" gorm:"uniqueIndex;type:uuid"`
	Title     string    `json:"title"`
	Number    float32   `json:"number"`
	Pages     []Page    `json:"pages"`
	CreatedAt time.Time `json:"createdAt"`
}

func (chapter *Chapter) BeforeCreate(*gorm.DB) error {
	if chapter.ChapterId == uuid.Nil {
		chapter.ChapterId = uuid.New()
	}
	chapter.CreatedAt = time.Now()
	return nil
}

type Page struct {
	PageId    uuid.UUID `json:"pageId" gorm:"primaryKey;type:uuid"`
	MangaId   uuid.UUID `json:"mangaId"`
	ChapterId uuid.UUID `json:"chapterId"`
	Number    uint32    `json:"number"`
	Image     string    `json:"image"`
}

func (page *Page) BeforeCreate(*gorm.DB) error {
	if page.PageId == uuid.Nil {
		page.PageId = uuid.New()
	}
	return nil
}
