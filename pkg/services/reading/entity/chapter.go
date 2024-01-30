package entity

import "time"

type Chapter struct {
	ChapterId string
	Number    float32
	Title     string
	Pages     uint32
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
