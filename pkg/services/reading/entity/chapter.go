package entity

import (
	"time"
)

type Chapter struct {
	ChapterId string
	MangaId   string
	Number    float32
	Title     string
	Pages     []Page
	CreatedAt time.Time
}

type ChapterHead struct {
	ChapterId string
	Number    float32
	Title     string
	Pages     uint32
	Read      bool
	CreatedAt time.Time
}

type Page struct {
	PageId    string
	MangaId   string
	ChapterId string
	Number    uint32
	Image     string
}
