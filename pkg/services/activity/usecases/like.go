package usecases

import (
	"errors"
	commonerrors "micromango/pkg/common/errors"
	"micromango/pkg/services/activity/entity"
)

// Like is the case when a user likes or dislikes a manga.
type Like struct {
	LikeRepository LikeRepository
	Catalog        CatalogClient
}

func (l *Like) Like(userId, mangaId string) error {
	r := entity.LikeRecord{
		UserId:  userId,
		MangaId: mangaId,
	}
	if err := l.LikeRepository.Save(r); err != nil {
		return err
	}
	likes, err := l.LikeRepository.CountByMangaId(mangaId)
	if err != nil {
		return err
	}
	return l.Catalog.SetLikesNumber(mangaId, likes)
}

func (l *Like) HasLike(userId, mangaId string) (bool, error) {
	r := entity.LikeRecord{
		UserId:  userId,
		MangaId: mangaId,
	}
	_, err := l.LikeRepository.Find(r)
	if err != nil {
		if errors.Is(err, &commonerrors.ErrNotFound{}) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (l *Like) Dislike(userId, mangaId string) error {
	r := entity.LikeRecord{
		UserId:  userId,
		MangaId: mangaId,
	}
	return l.LikeRepository.Remove(r)
}
