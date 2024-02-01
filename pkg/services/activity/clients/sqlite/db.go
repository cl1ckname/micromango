package sqlite

import (
	"errors"
	"gorm.io/gorm"
	commonerrors "micromango/pkg/common/errors"
	"micromango/pkg/services/activity/entity"
)

type DB struct {
	db *gorm.DB
}

func (d *DB) Save(record entity.LikeRecord) error {
	return d.db.Save(&record).Error
}

func (d *DB) Remove(record entity.LikeRecord) error {
	if err := d.db.Delete(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return commonerrors.ThrowNotFound(err.Error())
		}
		return err
	}
	return nil
}

func (d *DB) Find(record entity.LikeRecord) (entity.LikeRecord, error) {
	if err := d.db.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.LikeRecord{}, commonerrors.ThrowNotFound(err.Error())
		}
		return entity.LikeRecord{}, err
	}
	return record, nil
}

func (d *DB) CountByMangaId(mangaId string) (res uint64, err error) {
	err = d.db.Raw(`SELECT count(*) FROM like_records WHERE manga_id = ?`, mangaId).Scan(&res).Error
	return res, err
}
