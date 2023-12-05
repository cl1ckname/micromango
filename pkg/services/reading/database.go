package reading

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/reading"
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
	if err := db.AutoMigrate(&Chapter{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Page{}); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getMangaContent(db *gorm.DB, mangaId string) (m []Chapter, err error) {
	mangaUUID, err := uuid.Parse(mangaId)
	if err != nil {
		return nil, err
	}
	err = db.Find(&m, &Chapter{MangaId: mangaUUID}).Error
	return
}

func getChapter(db *gorm.DB, chapterId string) (c Chapter, err error) {
	c.ChapterId, err = uuid.Parse(chapterId)
	if err != nil {
		return Chapter{}, err
	}
	err = db.Model(&c).Preload("Pages").First(&c).Error
	return
}

func addChapter(db *gorm.DB, req *pb.AddChapterRequest) (c Chapter, err error) {
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return Chapter{}, err
	}
	c.Title = req.Title
	c.MangaId = mangaId
	c.Number = req.Number
	err = db.Create(&c).Error
	return
}

func updateChapter(db *gorm.DB, req *pb.UpdateChapterRequest) (c Chapter, err error) {
	c.ChapterId, err = uuid.Parse(req.ChapterId)
	if err != nil {
		return
	}
	if err = db.First(&c).Error; err != nil {
		return
	}
	c.Title = utils.DerefOrDefault(req.Title, c.Title)
	c.Number = utils.DerefOrDefault(req.Number, c.Number)
	err = db.Save(&c).Error
	return
}

func getPage(db *gorm.DB, pageId string) (p Page, err error) {
	p.PageId, err = uuid.Parse(pageId)
	if err != nil {
		return Page{}, err
	}
	err = db.First(&p).Error
	return p, err
}

func addPage(db *gorm.DB, req Page) (Page, error) {
	err := db.Create(&req).Error
	return req, err
}
