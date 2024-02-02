package entity

import (
	"time"
)

// Rate is a record of a user's rating of a manga.
type Rate struct {
	MangaId   string
	UserId    string
	Rate      uint32
	CreatedAt time.Time
}
