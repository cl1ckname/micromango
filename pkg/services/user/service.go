package user

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/user"
)

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)
	serv := service{
		db:        database,
		salt:      c.Salt,
		jwtSecret: c.JwtSecret,
	}
	baseServer := grpc.NewServer()
	pb.RegisterUserServer(baseServer, &serv)
	return common.StartGrpcService(ctx, c.Addr, baseServer)
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

func (s *service) Auth(_ context.Context, req *pb.AuthRequest) (*pb.UserResponse, error) {
	claims, err := s.auth(req.Token)
	if err != nil {
		return nil, err
	}
	userId := claims.UserId
	u, err := findById(s.db, userId)
	if err != nil {
		return nil, err
	}
	return u.ToPb(), nil
}

func (s *service) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	u, err := findById(s.db, req.UserId)
	if err != nil {
		return nil, err
	}
	return u.ToPb(), nil
}
