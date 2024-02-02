package usecases

import (
	"micromango/pkg/services/activity/entity"
)

type LikeRepository interface {
	SaveLike(entity.Like) error
	Remove(entity.Like) error
	FindLikeRecord(entity.Like) (entity.Like, error)
	CountByMangaId(string) (uint64, error)
}

type CatalogClient interface {
	SetLikesNumber(string, uint64) error
	SetAvgRate(string, float32, uint64) error
}

type RateRepository interface {
	SaveRate(like entity.Rate) error
	AvgRate(mangaId string) (float32, uint64, error)
	GetRate(userId, mangaId string) (uint32, error)
	GetRateList(userId string, mangaId []string) ([]entity.Rate, error)
}

type ReadRepository interface {
	SaveReadRecord(entity.ReadRecord) error
	GetReadChapters(userId string, mangaId string) (chapterIds []string, err error)
}
