package tests

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "micromango/pkg/grpc/user"
	"micromango/pkg/services/user"
	"testing"
)

const testServerAddr = ":50012"

func TestService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	salt := "salt"
	jwtSecret := "jwt-secret"
	user.Run(ctx, user.Config{
		Addr:      testServerAddr,
		DbAddr:    ":memory:",
		Salt:      salt,
		JwtSecret: jwtSecret,
	})

	conn, err := grpc.Dial(testServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	cc := pb.NewUserClient(conn)

	res, err := cc.Register(context.TODO(), &pb.RegisterRequest{
		Username: "clickname",
		Email:    "click@name.ru",
		Password: "qwerty",
	})
	require.NoError(t, err)

	uid := res.UserId
	_, err = uuid.Parse(uid)
	require.NoError(t, err)

	expected := pb.UserResponse{
		UserId:   uid,
		Username: "clickname",
		Email:    "click@name.ru",
		Picture:  "",
	}
	expBytes, _ := json.Marshal(&expected)
	resBytes, _ := json.Marshal(res)
	require.Equal(t, string(expBytes), string(resBytes))

	require.True(t, proto.Equal(&expected, res))

	loginRes, err := cc.Login(context.TODO(), &pb.LoginRequest{
		Email:    "click@name.ru",
		Password: "qwerty",
	})
	require.NoError(t, err)

	tokenS := loginRes.AccessToken
	var claims user.Claims
	token, err := jwt.ParseWithClaims(tokenS, &claims, func(*jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	require.NoError(t, err)
	require.True(t, token.Valid)
	require.Equal(t, uid, claims.UserId)
}
