package entity

import (
	"time"
)

// RateRecord is a record of a user's rating of a manga.
type RateRecord struct {
	MangaId   string
	UserId    string
	Rate      uint32
	CreatedAt time.Time
}
