package usecases

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	commonerr "micromango/pkg/common/errors"
	"micromango/pkg/common/utils"
	"micromango/pkg/services/reading/entity"
)

func (c *Case) GetMangaContent(mangaId string, userId *string) ([]entity.ChapterHead, error) {
	m, err := c.repo.GetContent(mangaId)
	if err != nil {
		if errors.Is(err, &commonerr.ErrNotFound{}) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("content for %s not found", mangaId))
		}
		return nil, err
	}

	readSet := make(map[string]struct{})
	if userId != nil {
		resp, err := c.activity.GetReadChapters(*userId, mangaId)
		if err != nil {
			return nil, err
		}
		for _, uid := range resp {
			readSet[uid] = struct{}{}
		}
	}

	resp := utils.Map(m, func(c entity.Chapter) entity.ChapterHead {
		_, read := readSet[c.ChapterId]
		return entity.ChapterHead{
			ChapterId: c.ChapterId,
			Number:    c.Number,
			Title:     c.Title,
			Pages:     c.Pages,
			Read:      read,
			CreatedAt: c.CreatedAt,
		}
	})
	return resp, nil
}
