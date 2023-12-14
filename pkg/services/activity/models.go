package activity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type LikeRecord struct {
	MangaId   uuid.UUID `gorm:"primaryKey; type:uuid"`
	UserId    uuid.UUID `gorm:"primaryKey; type:uuid"`
	CreatedAt time.Time
}

func (lr *LikeRecord) BeforeCreate(*gorm.DB) error {
	lr.CreatedAt = time.Now()
	return nil
}
