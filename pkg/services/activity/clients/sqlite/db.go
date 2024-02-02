package sqlite

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	commonerrors "micromango/pkg/common/errors"
	"micromango/pkg/common/utils"
	"micromango/pkg/services/activity/entity"
)

type DB struct {
	Db *gorm.DB
}

func (d *DB) SaveReadRecord(record entity.ReadRecord) error {
	rr, err := readRecordFromEntity(record)
	if err != nil {
		return err
	}
	return d.Db.Save(&rr).Error
}

func (d *DB) GetReadChapters(userId string, mangaId string) (chapterIds []string, err error) {
	mid, err := uuid.Parse(mangaId)
	if err != nil {
		return nil, err
	}
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	var r []ReadRecord
	rr := ReadRecord{
		MangaId: mid,
		UserId:  uid,
	}
	err = d.Db.Find(&r, &rr).Error
	chapterIds = utils.Map(r, func(r ReadRecord) string { return r.ChapterId.String() })
	return
}

func (d *DB) GetRate(userId, mangaId string) (uint32, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return 0, err
	}
	mid, err := uuid.Parse(mangaId)
	if err != nil {
		return 0, err
	}
	rr := RateRecord{
		MangaId: mid,
		UserId:  uid,
	}

	var res uint32
	err = d.Db.First(&rr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, commonerrors.ThrowNotFound("not found")
	}
	return res, err
}

func (d *DB) GetRateList(userId string, mangaId []string) ([]entity.Rate, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	var resp []RateRecord
	err = d.Db.Table("rate_records").
		Where("user_id = ? and manga_id in (?)", uid, mangaId).
		Scan(&resp).Error
	if err != nil {
		return nil, err
	}
	rates := utils.Map(resp, rateToEntity)
	return rates, err
}

func (d *DB) SaveRate(rate entity.Rate) error {
	model, err := rateFromEntity(rate)
	if err != nil {
		return err
	}
	return d.Db.Save(model).Error
}

func (d *DB) AvgRate(mangaId string) (float32, uint64, error) {
	var res AvgRateEntry
	err := d.Db.Raw(selectAvgRateQuery, mangaId).Scan(&res).Error
	return res.Rate, res.Voters, err
}

func (d *DB) Find(record entity.Like) (entity.Like, error) {
	likeModel, err := likeFromEntity(record)
	if err != nil {
		return entity.Like{}, err
	}

	if err := d.Db.First(&likeModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Like{}, commonerrors.ThrowNotFound(err.Error())
		}
		return entity.Like{}, err
	}

	return likeToEntity(likeModel), nil
}

func (d *DB) CountByMangaId(mangaId string) (res uint64, err error) {
	err = d.Db.Raw(`SELECT count(*) FROM like_records WHERE manga_id = ?`, mangaId).Scan(&res).Error
	return res, err
}
