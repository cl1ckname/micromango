package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"micromango/pkg/grpc/user"
	"net/http"
)

func (s *server) Register(ctx echo.Context) error {
	var registerReq user.RegisterRequest
	if err := ctx.Bind(&registerReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.user.Register(context.TODO(), &registerReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusCreated, resp)
}

func (s *server) Login(ctx echo.Context) error {
	var loginReq user.LoginRequest
	if err := ctx.Bind(&loginReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.user.Login(context.TODO(), &loginReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) GetUser(ctx echo.Context) error {
	var getUserReq user.GetUserRequest
	if err := ctx.Bind(&getUserReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.user.GetUser(context.TODO(), &getUserReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}