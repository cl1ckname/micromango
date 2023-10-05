package catalog

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	pb "micromango/pkg/grpc/catalog"
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

func AddManga(db *gorm.DB, req *pb.AddMangaRequest) (Manga, error) {
	m := Manga{
		Title:       req.Title,
		Cover:       req.Cover,
		Description: req.Description,
	}
	if res := db.Create(m); res.Error != nil {
		return Manga{}, res.Error
	}
	return m, nil
}
