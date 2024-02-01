package entity

import (
	"time"
)

// LikeRecord is record about a user like a manga.
type LikeRecord struct {
	MangaId   string
	UserId    string
	CreatedAt time.Time
}
