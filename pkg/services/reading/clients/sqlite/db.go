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

func (d *DB) GetChapter(chapterId string) (e entity.Chapter, err error) {
	var c Chapter
	c.ChapterId, err = uuid.Parse(chapterId)
	if err != nil {
		return entity.Chapter{}, err
	}
	err = d.db.Model(&c).Preload("Pages").First(&c).Error
	if err != nil {
		return entity.Chapter{}, err
	}
	e = ChapterToEntity(c)
	return
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
	m = utils.Map(chapters, ChapterToEntity)
	return
}

func (d *DB) SaveChapter(chapter entity.Chapter) (entity.Chapter, error) {
	c, err := ChapterFromEntity(chapter)
	if err != nil {
		return entity.Chapter{}, err
	}
	err = d.db.Save(&c).Error
	if err != nil {
		return entity.Chapter{}, err
	}
	return ChapterToEntity(c), nil
}

func (d *DB) SavePage(page entity.Page) (entity.Page, error) {
	pageModel, err := PageFromEntity(page)
	if err != nil {
		return entity.Page{}, err
	}
	err = d.db.Save(&pageModel).Error
	if err != nil {
		return entity.Page{}, err
	}
	return entity.Page{}, err
}

func (d *DB) GetPage(pageId string) (entity.Page, error) {
	pageUUID, err := uuid.Parse(pageId)
	if err != nil {
		return entity.Page{}, err
	}
	var pageModel Page
	pageModel.PageId = pageUUID
	err = d.db.Model(&pageModel).First(&pageModel).Error
	if err != nil {
		return entity.Page{}, err
	}
	return PageToEntity(pageModel), nil
}
