package usecases

import (
	"micromango/pkg/services/activity/entity"
)

type Read struct {
	ReadRepository ReadRepository
}

func (r *Read) ReadChapter(rr entity.ReadRecord) error {
	return r.ReadRepository.SaveReadRecord(rr)
}

func (r *Read) GetReadChapters(userId, mangaId string) ([]string, error) {
	return r.ReadRepository.GetReadChapters(userId, mangaId)
}
