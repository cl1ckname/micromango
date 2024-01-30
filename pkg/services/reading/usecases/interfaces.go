package usecases

import "micromango/pkg/services/reading/entity"

type MangaRepository interface {
	GetContent(mangaId string) ([]entity.Chapter, error)
}

type ActivityService interface {
	GetReadChapters(userId, mangaId string) ([]string, error)
}
