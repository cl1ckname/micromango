package user

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	jwt.Claims
	UserId string `json:"userId"`
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

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, &Claims{
		Claims: jwt.RegisteredClaims{
			Issuer:    userModel.UserId.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UserId: userModel.UserId.String(),
	})
	return token, nil
}
