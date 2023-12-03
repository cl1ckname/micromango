package user

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"micromango/pkg/common"
	"micromango/pkg/grpc/profile"
	"time"
)

func (s *service) login(email string, password string) (*jwt.Token, error) {
	passwordHash := hashString(password, s.salt)

	userModel := User{
		PasswordHash: passwordHash,
		Email:        email,
	}
	err := s.db.First(&userModel).Error
	if err != nil {
		return nil, err
	}

	p, err := s.profile.Get(context.TODO(), &profile.GetRequest{UserId: userModel.UserId.String()})
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, common.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userModel.UserId.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UserId:   userModel.UserId.String(),
		Username: p.Username,
	})
	return token, nil
}

func (s *service) auth(token string) (common.Claims, error) {
	t, err := jwt.ParseWithClaims(token, &common.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return common.Claims{}, err
	}
	if !t.Valid {
		return common.Claims{}, fmt.Errorf("invalid token")
	}
	claims, ok := t.Claims.(*common.Claims)
	if !ok {
		return common.Claims{}, fmt.Errorf("invalid claims")
	}
	return *claims, nil
}
