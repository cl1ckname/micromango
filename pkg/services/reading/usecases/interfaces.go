package usecases

import (
	"micromango/pkg/common"
	"micromango/pkg/services/reading/entity"
)

type ChapterRepository interface {
	GetContent(mangaId string) ([]entity.Chapter, error)
	GetChapter(chapterId string) (entity.Chapter, error)
	SaveChapter(chapter entity.Chapter) (entity.Chapter, error)
}

type ActivityService interface {
	GetReadChapters(userId, mangaId string) ([]string, error)
}

type PageRepository interface {
	SavePage(page entity.Page) (entity.Page, error)
	GetPage(pageId string) (entity.Page, error)
}

type StaticService interface {
	UploadPage(mangaId, chapterId string, image common.File) (string, error)
}
