package mock

import (
	commonerr "micromango/pkg/common/errors"
	"micromango/pkg/services/reading/entity"
)

var testChapters = map[string]entity.Chapter{
	"1": {
		ChapterId: "1",
		MangaId:   "1",
		Number:    1,
		Title:     "title",
		Pages:     []entity.Page{{PageId: "1", MangaId: "1", ChapterId: "1", Number: 1}},
	},
	"2": {
		ChapterId: "2",
		MangaId:   "1",
		Number:    2,
		Title:     "title2",
		Pages:     []entity.Page{{PageId: "2", MangaId: "1", ChapterId: "1", Number: 2}},
	},
	"3": {
		ChapterId: "3",
		MangaId:   "1",
		Number:    3,
		Title:     "title3",
		Pages:     []entity.Page{{PageId: "3", MangaId: "1", ChapterId: "1", Number: 3}},
	},
}

// ChapterRepository is a mock implementation of ChapterRepository.
type ChapterRepository struct{}

func (c *ChapterRepository) GetContent(_ string) ([]entity.Chapter, error) {
	var chapters []entity.Chapter
	for _, chapter := range []string{"1", "2", "3"} {
		chapters = append(chapters, testChapters[chapter])
	}
	return chapters, nil
}

func (c *ChapterRepository) GetChapter(chapterId string) (entity.Chapter, error) {
	chapter, ok := testChapters[chapterId]
	if !ok {
		return entity.Chapter{}, &commonerr.ErrNotFound{}
	}
	return chapter, nil
}

func (c *ChapterRepository) SaveChapter(chapter entity.Chapter) (entity.Chapter, error) {
	chapter.ChapterId = "1"
	return chapter, nil
}
