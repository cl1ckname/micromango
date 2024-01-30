package sqlite

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"micromango/pkg/common/utils"
	"micromango/pkg/services/reading/entity"
	"os"
	"time"
)

type DB struct {
	db *gorm.DB
}

func Connect(connStr string) *DB {
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
	db, err := gorm.Open(sqlite.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err := db.AutoMigrate(&Chapter{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Page{}); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db}
}

const GetContentQuery = `SELECT c.chapter_id, c.number, title, count(p.chapter_id) as pages, created_at  
    FROM chapters c LEFT JOIN pages p on c.chapter_id = p.chapter_id and c.manga_id = p.manga_id
    WHERE c.manga_id = ? GROUP BY c.chapter_id`

func (d *DB) GetContent(mangaId string) (m []entity.Chapter, err error) {
	mangaUUID, err := uuid.Parse(mangaId)
	if err != nil {
		return nil, err
	}
	var chapters []Chapter
	if err = d.db.Raw(GetContentQuery, mangaUUID).Scan(&chapters).Error; err != nil {
		return nil, err
	}
	m = utils.Map(chapters, func(c Chapter) entity.Chapter {
		return entity.Chapter{
			ChapterId: c.ChapterId.String(),
			Number:    c.Number,
			Title:     c.Title,
			Pages:     uint32(len(c.Pages)),
			CreatedAt: c.CreatedAt,
		}
	})
	return
}
