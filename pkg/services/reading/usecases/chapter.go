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

type Chapter struct {
	Repository ChapterRepository
	Activity   ActivityService
}

func (c *Chapter) GetMangaContent(mangaId string, userId *string) ([]entity.ChapterHead, error) {
	m, err := c.Repository.GetContent(mangaId)
	if err != nil {
		if errors.Is(err, &commonerr.ErrNotFound{}) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("content for %s not found", mangaId))
		}
		return nil, err
	}

	readSet := make(map[string]struct{})
	if userId != nil {
		resp, err := c.Activity.GetReadChapters(*userId, mangaId)
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
			Pages:     uint32(len(c.Pages)),
			Read:      read,
			CreatedAt: c.CreatedAt,
		}
	})
	return resp, nil
}

func (c *Chapter) GetChapter(chapterId string) (entity.Chapter, error) {
	return c.Repository.GetChapter(chapterId)
}

func (c *Chapter) AddChapter(chapter entity.Chapter) (entity.Chapter, error) {
	return c.Repository.SaveChapter(chapter)
}

func (c *Chapter) UpdateChapter(chapterId string, data entity.UpdateChapterDto) (entity.Chapter, error) {
	chapterToUpdate, err := c.Repository.GetChapter(chapterId)
	if err != nil {
		return entity.Chapter{}, err
	}
	chapterToUpdate.Title = utils.DerefOrDefault(data.Title, chapterToUpdate.Title)
	chapterToUpdate.Number = utils.DerefOrDefault(data.Number, chapterToUpdate.Number)
	return c.Repository.SaveChapter(chapterToUpdate)
}
