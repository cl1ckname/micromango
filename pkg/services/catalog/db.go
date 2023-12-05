package catalog

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
	if err := db.AutoMigrate(&Manga{}); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetManga(db *gorm.DB, mangaId string) (Manga, error) {
	mangaUUID, err := uuid.Parse(mangaId)
	if err != nil {
		return Manga{}, err
	}
	m := Manga{MangaId: mangaUUID}
	if res := db.First(&m); res.Error != nil {
		return Manga{}, res.Error
	}
	return m, nil
}

func GetMangas(db *gorm.DB) ([]Manga, error) {
	var mangas []Manga
	if res := db.Find(&mangas); res.Error != nil {
		return nil, res.Error
	}
	return mangas, nil
}

func AddManga(db *gorm.DB, m Manga) (Manga, error) {
	if res := db.Create(&m); res.Error != nil {
		return Manga{}, res.Error
	}
	return m, nil
}

func SaveManga(db *gorm.DB, m Manga) (Manga, error) {
	if res := db.Save(&m); res.Error != nil {
		return Manga{}, res.Error
	}
	return m, nil
}

func DeleteManga(db *gorm.DB, mangaId string) error {
	mangaUuid, err := uuid.Parse(mangaId)
	if err != nil {
		return err
	}
	toDelete := Manga{MangaId: mangaUuid}
	return db.Delete(&toDelete).Error
}

func GetMany(db *gorm.DB, listId []string) (m []Manga, err error) {
	err = db.Find(&m, listId).Error
	return
}
