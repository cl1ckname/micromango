package sqlite

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RateRecord struct {
	MangaId   uuid.UUID `gorm:"primaryKey; type:uuid"`
	UserId    uuid.UUID `gorm:"primaryKey; type:uuid"`
	Rate      uint32
	CreatedAt time.Time
}

func (rr *RateRecord) BeforeCreate(*gorm.DB) error {
	rr.CreatedAt = time.Now()
	return nil
}

type ReadRecord struct {
	MangaId   uuid.UUID `gorm:"primaryKey; type:uuid" json:"mangaId"`
	ChapterId uuid.UUID `gorm:"primaryKey; type:uuid" json:"chapterId"`
	UserId    uuid.UUID `gorm:"primaryKey; type:uuid" json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

func (rr *ReadRecord) BeforeCreate(*gorm.DB) error {
	rr.CreatedAt = time.Now()
	return nil
}

type LikeRecord struct {
	MangaId   uuid.UUID `gorm:"primaryKey; type:uuid"`
	UserId    uuid.UUID `gorm:"primaryKey; type:uuid"`
	CreatedAt time.Time
}

func (lr *LikeRecord) BeforeCreate(*gorm.DB) error {
	lr.CreatedAt = time.Now()
	return nil
}

type AvgRateEntry struct {
	Rate   float32
	Voters uint64
}
