package entity

import "micromango/pkg/common"

type UpdateChapterDto struct {
	Title  *string  `json:"title"`
	Number *float32 `json:"number"`
}

type AddPageDto struct {
	MangaId   string       `json:"manga"`
	ChapterId string       `json:"chapter"`
	Number    uint32       `json:"number"`
	Image     *common.File `json:"image"`
}
