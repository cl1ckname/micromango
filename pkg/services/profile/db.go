package profile

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
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open(addr), &gorm.Config{
		Logger: newLogger,
	})
	if err := db.AutoMigrate(&Profile{}); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateProfile(db *gorm.DB, p Profile) (Profile, error) {
	err := db.Create(&p).Error
	return p, err
}

func SaveProfile(db *gorm.DB, p Profile) (Profile, error) {
	err := db.Save(&p).Error
	return p, err
}

func FindOne(db *gorm.DB, userId uuid.UUID) (p Profile, err error) {
	p.UserId = userId
	err = db.First(&p).Error
	return
}
