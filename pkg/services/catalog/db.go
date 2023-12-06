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
	var args []interface{}
	sql := `SELECT * FROM mangas m `
	if l := len(include); l != 0 {
		sql += `where exists(
		select * from manga_genres mg
		where m.manga_id = mg.manga_manga_id and
				mg.genre_genre_id in (?)
		group by manga_manga_id
		having count(mg.genre_genre_id) = ?)`
		args = append(args, include, l)
	}
	if len(exclude) != 0 {
		sql += `
		intersect
		select * from mangas m where exists (
		select manga_manga_id, sum(case when genre_genre_id in (?) then 1 else 0 end) as exclude from manga_genres
		where manga_manga_id == m.manga_id
		group by manga_manga_id having exclude == 0)`
		args = append(args, exclude)
	}

	err = db.Raw(sql, args...).Scan(&m).Error
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
