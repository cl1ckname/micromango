package mock

import (
	"micromango/pkg/common/errors"
	"micromango/pkg/services/reading/entity"
)

type PageRepository struct{}

func (p *PageRepository) GetPage(pageId string) (entity.Page, error) {
	if pageId == "0" {
		return entity.Page{}, errors.ThrowNotFound("not found")
	}
	return entity.Page{PageId: pageId, Number: 1, ChapterId: "1", MangaId: "1"}, nil
}

func (p *PageRepository) SavePage(page entity.Page) (entity.Page, error) {
	page.PageId = "0"
	return page, nil
}
