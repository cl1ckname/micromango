package activity

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Connect(addr string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open(addr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&LikeRecord{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&RateRecord{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&ReadRecord{}); err != nil {
		panic(err)
	}
	return db
}

func SaveLike(db *gorm.DB, mangaId, userId uuid.UUID) (LikeRecord, error) {
	lr := LikeRecord{
		MangaId: mangaId,
		UserId:  userId,
	}
	err := db.Create(&lr).Error
	return lr, err
}

func RemoveLike(db *gorm.DB, mangaId, userId uuid.UUID) error {
	lr := LikeRecord{
		MangaId: mangaId,
		UserId:  userId,
	}
	return db.Delete(&lr).Error
}

func LikesCount(db *gorm.DB, mangaId uuid.UUID) (uint64, error) {
	var res uint64
	err := db.Raw(`SELECT count(*) FROM like_records WHERE manga_id = ?`, mangaId).Scan(&res).Error
	return res, err
}

func HasLike(db *gorm.DB, userId uuid.UUID, mangaId uuid.UUID) (bool, error) {
	lr := LikeRecord{
		MangaId: mangaId,
		UserId:  userId,
	}
	err := db.First(&lr).Error
	if err == nil {
		return true, nil
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, err
}

func SaveRate(db *gorm.DB, userId, mangaId uuid.UUID, rate uint32) error {
	rr := RateRecord{
		MangaId: mangaId,
		UserId:  userId,
		Rate:    rate,
	}
	return db.Save(&rr).Error
}

func AvgRate(db *gorm.DB, mangaId uuid.UUID) (AvgRateEntry, error) {
	var res AvgRateEntry
	sql := `SELECT AVG(rate) as rate, COUNT(*) as voters FROM rate_records WHERE manga_id = ?`
	err := db.Raw(sql, mangaId).Scan(&res).Error
	return res, err
}

func UserRate(db *gorm.DB, userId, mangaId uuid.UUID) (*float32, error) {
	rr := RateRecord{
		MangaId: mangaId,
		UserId:  userId,
	}
	var res float32
	err := db.First(&rr).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	res = float32(rr.Rate)
	return &res, err
}

func RateList(db *gorm.DB, userId uuid.UUID, mangaIds []string) ([]RateRecord, error) {
	var resp []RateRecord
	err := db.Table("rate_records").
		Where("user_id = ? and manga_id in (?)", userId, mangaIds).
		Scan(&resp).Error
	return resp, err
}

func ReadChapter(db *gorm.DB, userId, mangaId, chapterId uuid.UUID) error {
	rr := ReadRecord{
		MangaId:   mangaId,
		ChapterId: chapterId,
		UserId:    userId,
	}
	return db.Save(&rr).Error
}

func ReadChapters(db *gorm.DB, userId, mangaId uuid.UUID) (r []uuid.UUID, err error) {
	rr := RateRecord{
		MangaId: mangaId,
		UserId:  userId,
	}
	err = db.Find(&r, &rr).Error
	return
}
