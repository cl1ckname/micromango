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
}
