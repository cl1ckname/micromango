package user

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/user"
)

func Run() {
	database := Connect("reading.sqlite")
	serv := service{
		db:        database,
		salt:      "qwerty",
		jwtSecret: "generateLater",
	}
	addr := fmt.Sprintf(":%d", 50002)
	if err := common.RunGRPCServer(addr, func(registrar grpc.ServiceRegistrar) {
		pb.RegisterUserServer(registrar, &serv)
	}); err != nil {
		panic(err)
	}
}

type service struct {
	pb.UnimplementedUserServer
	db        *gorm.DB
	salt      string
	jwtSecret string
}

func (s *service) Register(_ context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	u, err := findByEmail(s.db, req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	u.Username = req.Username
	u.Email = req.Email
	u.PasswordHash = hashString(req.Password, s.salt)
	u, err = saveUser(s.db, u)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId:   u.UserId.String(),
		Username: u.Username,
		Email:    u.Email,
		Picture:  u.Picture,
	}, nil
}

func (s *service) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, err := findByEmail(s.db, req.Email)
	if err != nil {
		return nil, err
	}
	token, err := s.login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	accessToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}
	accessTokenExpired, err := token.Claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:        accessToken,
		AccessTokenExpired: timestamppb.New(accessTokenExpired.Time),
	}, nil
}
