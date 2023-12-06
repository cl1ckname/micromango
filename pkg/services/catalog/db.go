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
	if err := db.AutoMigrate(&Manga{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Genre{}); err != nil {
		panic(err)
	}
	return db
}

func GetManga(db *gorm.DB, mangaId string) (m Manga, err error) {
	mangaUUID, err := uuid.Parse(mangaId)
	if err != nil {
		return Manga{}, err
	}
	m = Manga{MangaId: mangaUUID}
	if res := db.Preload("Genres").First(&m); res.Error != nil {
		return Manga{}, res.Error
	}
	return m, nil
}

func GetMangas(db *gorm.DB, include []uint32, exclude []uint32) (m []Manga, err error) {
	err = db.
		Joins("INNER JOIN manga_genres ON manga_genres.manga_manga_id = "+
			"mangas.manga_id AND manga_genres.genre_genre_id in (?) AND manga_genres.genre_genre_id not in (?)", include, exclude).
		Find(&m).Error
	return
}

func AddManga(db *gorm.DB, m Manga) (Manga, error) {
	genres := m.Genres
	m.Genres = []Genre{}
	if res := db.Create(&m); res.Error != nil {
		return Manga{}, res.Error
	}
	if err := db.Model(&m).Association("Genres").Append(genres); err != nil {
		return Manga{}, err
	}
	return GetManga(db, m.MangaId.String())
}

func SaveManga(db *gorm.DB, m Manga) (Manga, error) {
	genres := m.Genres
	m.Genres = []Genre{}
	if res := db.Save(&m); res.Error != nil {
		return Manga{}, res.Error
	}
	if err := db.Model(&m).Association("Genres").Replace(genres); err != nil {
		return Manga{}, err
	}
	return GetManga(db, m.MangaId.String())
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
