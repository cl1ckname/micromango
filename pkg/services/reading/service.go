package reading

import (
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/catalog/db"
)

type service struct {
	pb.UnimplementedReadingServer
	db *gorm.DB
}

func Run() {
	database := db.Connect("catalog.sqlite")
	serv := service{
		db: database,
	}
	addr := fmt.Sprintf(":%d", 50001)
	if err := common.RunGRPCServer(addr, func(registrar grpc.ServiceRegistrar) {
		pb.RegisterReadingServer(registrar, &serv)
	}); err != nil {
		panic(err)
	}
}
