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
