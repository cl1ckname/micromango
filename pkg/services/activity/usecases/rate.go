package usecases

type Rate struct {
	Repository    RateRepository
	CatalogClient CatalogClient
}

func (r *Rate) RateManga(userId, mangaId string, rate uint32) error {
	if err := r.Repository.Save(userId, mangaId, rate); err != nil {
		return err
	}
	avg, voters, err := r.Repository.AvgRate(mangaId)
	if err != nil {
		return err
	}
	return r.CatalogClient.SetAvgRate(mangaId, avg, voters)
}

func (r *Rate) GetUserRate(userId, mangaId string) (uint32, error) {
	return r.Repository.Get(userId, mangaId)
}

func (r *Rate) GetRateList(userId string, ids []string) (map[string]uint32, error) {
	rates, err := r.Repository.GetList(userId, ids)
	if err != nil {
		return nil, err
	}
	ratesMap := make(map[string]uint32, len(rates))
	for _, rate := range rates {
		ratesMap[rate.MangaId] = rate.Rate
	}
	return ratesMap, nil
}
