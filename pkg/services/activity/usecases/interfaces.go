package usecases

import (
	"micromango/pkg/services/activity/entity"
)

type LikeRepository interface {
	Save(entity.LikeRecord) error
	Remove(entity.LikeRecord) error
	Find(entity.LikeRecord) (entity.LikeRecord, error)
	CountByMangaId(string) (uint64, error)
}

type CatalogClient interface {
	SetLikesNumber(string, uint64) error
	SetAvgRate(string, float32, uint64) error
}

type RateRepository interface {
	Save(userId, mangaId string, rate uint32) error
	AvgRate(mangaId string) (float32, uint64, error)
	Get(userId, mangaId string) (uint32, error)
	GetList(userId string, mangaId []string) ([]entity.Rate, error)
}

type ReadRepository interface {
	Save(entity.ReadRecord) error
	GetReadChapters(userId string, mangaId string) (chapterIds []string, err error)
}
