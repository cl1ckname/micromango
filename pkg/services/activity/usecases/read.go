package usecases

import (
	"micromango/pkg/services/activity/entity"
)

type Read struct {
	Repository ReadRepository
}

func (r *Read) ReadChapter(rr entity.ReadRecord) error {
	return r.Repository.SaveReadRecord(rr)
}

func (r *Read) GetReadChapters(userId, mangaId string) ([]string, error) {
	return r.Repository.GetReadChapters(userId, mangaId)
}
