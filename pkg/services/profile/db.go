package profile

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/share"
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
	if err := db.AutoMigrate(&ListRecord{}); err != nil {
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

func GetList(db *gorm.DB, userId uuid.UUID, list share.ListName) ([]ListRecord, error) {
	var lr []ListRecord
	cond := ListRecord{
		UserId:   userId,
		ListName: list,
	}
	err := db.Find(&lr, cond).Error
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
		UserId:  userUUID,
		MangaId: mangaUUID,
	}
	return db.Delete(&lr).Error
}

func FindListRecord(db *gorm.DB, req *pb.IsInListRequest) (*pb.IsInListResponse, error) {
	userUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	mangaUUID, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}

	lr := ListRecord{UserId: userUUID, MangaId: mangaUUID}
	err = db.First(&lr).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		return &pb.IsInListResponse{}, nil
	}
	return &pb.IsInListResponse{
		In:        utils.Ptr(lr.ListName),
		Timestamp: lr.CreatedAt.String(),
	}, nil
}

func ListStats(db *gorm.DB, mangaId uuid.UUID) (map[uint32]uint64, error) {
	var resSlice []ListStat
	sql := `SELECT list_name, count(user_id) as amount FROM list_records WHERE manga_id = ? GROUP BY manga_id, list_name`
	if err := db.Raw(sql, mangaId).Scan(&resSlice).Error; err != nil {
		return nil, err
	}
	res := make(map[uint32]uint64)
	for _, v := range resSlice {
		res[v.ListName] = v.Amount
	}
	return res, nil
}
