package entity

import (
	"time"
)

// Like is record about a user like a manga.
type Like struct {
	MangaId   string
	UserId    string
	CreatedAt time.Time
}
