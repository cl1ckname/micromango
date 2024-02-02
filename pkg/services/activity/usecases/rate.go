package usecases

import (
	"micromango/pkg/services/activity/entity"
)

type Rate struct {
	RateRepository RateRepository
	Catalog        CatalogClient
}

func (r *Rate) RateManga(like entity.Rate) error {
	if err := r.RateRepository.SaveRate(like); err != nil {
		return err
	}
	avg, voters, err := r.RateRepository.AvgRate(like.MangaId)
	if err != nil {
		return err
	}
	return r.Catalog.SetAvgRate(like.MangaId, avg, voters)
}

func (r *Rate) GetUserRate(userId, mangaId string) (uint32, error) {
	return r.RateRepository.GetRate(userId, mangaId)
}

func (r *Rate) GetRateList(userId string, ids []string) (map[string]uint32, error) {
	rates, err := r.RateRepository.GetRateList(userId, ids)
	if err != nil {
		return nil, err
	}
	ratesMap := make(map[string]uint32, len(rates))
	for _, rate := range rates {
		ratesMap[rate.MangaId] = rate.Rate
	}
	return ratesMap, nil
}
