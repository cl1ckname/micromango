package entity

type UpdateChapterDto struct {
	Title  *string  `json:"title"`
	Number *float32 `json:"number"`
}
