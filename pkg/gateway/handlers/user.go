package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/user"
	"net/http"
	"strings"
)

type userHandler struct {
	user user.UserClient
}

func RegisterUser(g *echo.Group, u user.UserClient) {
	h := userHandler{u}
	userGroup := g.Group("/user")

	userGroup.POST("/register", h.Register)
	userGroup.POST("/login", h.Login)
}

func (s *userHandler) Register(ctx echo.Context) error {
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

func (s *userHandler) Login(ctx echo.Context) error {
	var loginReq user.LoginRequest
	if err := ctx.Bind(&loginReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	resp, err := s.user.Login(context.TODO(), &loginReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	loginCookie := new(http.Cookie)
	loginCookie.Name = "auth"
	loginCookie.Value = resp.AccessToken
	loginCookie.HttpOnly = true
	loginCookie.Path = "/"
	ctx.SetCookie(loginCookie)

	publicCookie := new(http.Cookie)
	publicCookie.Name = "signin"
	tokenParts := strings.Split(resp.AccessToken, ".")
	if len(tokenParts) != 3 {
		log.Fatal("invalid token: ", resp.AccessToken)
	}
	publicCookie.Value = tokenParts[1]
	publicCookie.Path = "/"
	ctx.SetCookie(publicCookie)

	return ctx.JSON(http.StatusOK, resp)
}
