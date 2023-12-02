package profile

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	pb "micromango/pkg/grpc/profile"
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

func GetList(db *gorm.DB, req *pb.GetListRequest) ([]ListRecord, error) {
	userUuid, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, err
	}

	var lr []ListRecord
	cond := ListRecord{
		UserId:   userUuid,
		ListName: req.List,
	}
	err = db.Find(&lr, cond).Error
	return lr, err
}

func AddToList(db *gorm.DB, req *pb.AddToListRequest) error {
	userUUID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return err
	}
	mangaUUID, err := uuid.Parse(req.MangaId)
	if err != nil {
		return err
	}
	lr := ListRecord{
		UserId:   userUUID,
		MangaId:  mangaUUID,
		ListName: req.List,
	}
	return db.Save(&lr).Error
}

func RemoveFromList(db *gorm.DB, req *pb.RemoveFromListRequest) error {
	userUUID, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return err
	}
	mangaUUID, err := uuid.Parse(req.MangaId)
	if err != nil {
		return err
	}
	lr := ListRecord{
		UserId:   userUUID,
		MangaId:  mangaUUID,
		ListName: req.List,
	}
	return db.Delete(&lr).Error
}
