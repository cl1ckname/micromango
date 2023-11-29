package user

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userModel.UserId.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UserId:   userModel.UserId.String(),
		Username: userModel.Username,
	})
	return token, nil
}

func (s *service) auth(token string) (Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if !t.Valid {
		return Claims{}, fmt.Errorf("invalid token")
	}
	claims, ok := t.Claims.(*Claims)
	if !ok {
		return Claims{}, fmt.Errorf("invalid claims")
	}
	return *claims, nil
}
