package usecases

type Case struct {
	repo     MangaRepository
	activity ActivityService
}

func NewCase(repo MangaRepository, activity ActivityService) Case {
	return Case{repo, activity}
}
