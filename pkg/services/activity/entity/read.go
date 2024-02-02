package entity

import (
	"time"
)

type ReadRecord struct {
	MangaId   string
	ChapterId string
	UserId    string
	CreatedAt time.Time `json:"createdAt"`
}
