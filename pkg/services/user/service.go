package user

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"micromango/pkg/common"
	"micromango/pkg/grpc/profile"
	pb "micromango/pkg/grpc/user"
)

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)
	serv := service{
		db:        database,
		salt:      c.Salt,
		jwtSecret: c.JwtSecret,
	}

	conn, err := grpc.Dial(c.ProfileAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	serv.profile = profile.NewProfileClient(conn)

	baseServer := grpc.NewServer()
	pb.RegisterUserServer(baseServer, &serv)
	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	pb.UnimplementedUserServer
	db        *gorm.DB
	salt      string
	jwtSecret string
	profile   profile.ProfileClient
}

func (s *service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	u, err := findByEmail(s.db, req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	u.Email = req.Email
	u.PasswordHash = hashString(req.Password, s.salt)
	savedU, err := saveUser(s.db, u)
	if err != nil {
		return nil, err
	}
	if _, err := s.profile.Create(ctx, &profile.CreateRequest{
		UserId:   savedU.UserId.String(),
		Username: req.Username,
	}); err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId:   savedU.UserId.String(),
		Username: req.Username,
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
		fmt.Println("invalid token: ", err.Error())
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

func (s *service) Auth(_ context.Context, req *pb.AuthRequest) (*pb.UserResponse, error) {
	claims, err := s.auth(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserId:   claims.UserId,
		Username: claims.Username,
	}, nil
}
