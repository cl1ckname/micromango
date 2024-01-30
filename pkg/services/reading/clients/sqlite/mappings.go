package sqlite

import (
	"github.com/google/uuid"
	"micromango/pkg/common/utils"
	"micromango/pkg/services/reading/entity"
)

func ChapterToEntity(c Chapter) entity.Chapter {
	return entity.Chapter{
		ChapterId: c.ChapterId.String(),
		Number:    c.Number,
		Title:     c.Title,
		Pages:     utils.Map(c.Pages, PageToEntity),
		CreatedAt: c.CreatedAt,
	}
}

func PageToEntity(p Page) entity.Page {
	return entity.Page{
		PageId:    p.PageId.String(),
		MangaId:   p.MangaId.String(),
		ChapterId: p.ChapterId.String(),
		Number:    p.Number,
		Image:     p.Image,
	}
}

func ChapterFromEntity(c entity.Chapter) (Chapter, error) {
	chapterId, err := uuid.Parse(c.ChapterId)
	if err != nil {
		return Chapter{}, err
	}
	mangaId, err := uuid.Parse(c.MangaId)
	if err != nil {
		return Chapter{}, err
	}
	return Chapter{
		ChapterId: chapterId,
		MangaId:   mangaId,
		Title:     c.Title,
		Number:    c.Number,
		Pages:     nil,
		CreatedAt: c.CreatedAt,
	}, nil
}

func PageFromEntity(p entity.Page) (Page, error) {
	mangaId, err := uuid.Parse(p.MangaId)
	if err != nil {
		return Page{}, err
	}
	chapterId, err := uuid.Parse(p.ChapterId)
	if err != nil {
		return Page{}, err
	}
	return Page{
		MangaId:   mangaId,
		ChapterId: chapterId,
		Number:    p.Number,
		Image:     p.Image,
	}, nil
}
