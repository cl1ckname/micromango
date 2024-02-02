package sqlite

import (
	"github.com/google/uuid"
	"micromango/pkg/services/activity/entity"
)

func likeFromEntity(record entity.Like) (LikeRecord, error) {
	uid, err := uuid.Parse(record.UserId)
	if err != nil {
		return LikeRecord{}, err
	}
	mid, err := uuid.Parse(record.MangaId)
	if err != nil {
		return LikeRecord{}, err
	}
	likeModel := LikeRecord{
		MangaId: mid,
		UserId:  uid,
	}
	return likeModel, nil
}

func likeToEntity(record LikeRecord) entity.Like {
	return entity.Like{
		UserId:    record.UserId.String(),
		MangaId:   record.MangaId.String(),
		CreatedAt: record.CreatedAt,
	}
}

func rateFromEntity(record entity.Rate) (RateRecord, error) {
	uid, err := uuid.Parse(record.UserId)
	if err != nil {
		return RateRecord{}, err
	}
	mid, err := uuid.Parse(record.MangaId)
	if err != nil {
		return RateRecord{}, err
	}
	rateModel := RateRecord{
		MangaId: mid,
		UserId:  uid,
		Rate:    record.Rate,
	}
	return rateModel, nil
}

func rateToEntity(record RateRecord) entity.Rate {
	return entity.Rate{
		UserId:    record.UserId.String(),
		MangaId:   record.MangaId.String(),
		Rate:      record.Rate,
		CreatedAt: record.CreatedAt,
	}
}

func readRecordFromEntity(record entity.ReadRecord) (ReadRecord, error) {
	uid, err := uuid.Parse(record.UserId)
	if err != nil {
		return ReadRecord{}, err
	}
	mid, err := uuid.Parse(record.MangaId)
	if err != nil {
		return ReadRecord{}, err
	}
	cid, err := uuid.Parse(record.ChapterId)
	if err != nil {
		return ReadRecord{}, err
	}
	readModel := ReadRecord{
		MangaId:   mid,
		ChapterId: cid,
		UserId:    uid,
	}
	return readModel, nil
}
