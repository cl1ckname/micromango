package usecases

import (
	"micromango/pkg/services/reading/entity"
)

type Page struct {
	Repository PageRepository
	Static     StaticService
}

func (p *Page) GetPage(pageId string) (entity.Page, error) {
	return p.Repository.GetPage(pageId)
}

func (p *Page) AddPage(dto entity.AddPageDto) (res entity.Page, err error) {
	var imageUrl string
	if dto.Image != nil {
		imageUrl, err = p.Static.UploadPage(dto.ChapterId, dto.MangaId, *dto.Image)
		if err != nil {
			return entity.Page{}, err
		}
	}

	newPage, err := p.Repository.SavePage(entity.Page{
		MangaId:   dto.MangaId,
		ChapterId: dto.ChapterId,
		Number:    dto.Number,
		Image:     imageUrl,
	})
	if err != nil {
		return entity.Page{}, err
	}
	return newPage, nil
}
