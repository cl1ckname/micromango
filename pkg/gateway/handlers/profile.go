package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/share"
	"micromango/pkg/grpc/user"
	"net/http"
)

func RegisterProfile(e *echo.Group, p profile.ProfileClient) {
	handler := profileHandler{e, p}
	profileGroup := e.Group("/profile")

	profileGroup.GET("/:userId", handler.GetProfile)
	profileGroup.PUT("/:userId", handler.UpdateProfile)
	profileGroup.GET("/:userId/list", handler.GetList)
	profileGroup.POST("/:userId/list", handler.AddToList)
	profileGroup.DELETE("/:userId/list", handler.RemoveFromList)
}

type profileHandler struct {
	e       *echo.Group
	profile profile.ProfileClient
}

func (s *profileHandler) UpdateProfile(ctx echo.Context) error {
	var updateReq profile.UpdateRequest
	updateReq.UserId = ctx.Param("userId")
	if username := ctx.FormValue("username"); username != "" {
		updateReq.Username = utils.Ptr(username)
	}
	if bio := ctx.FormValue("bio"); bio != "" {
		updateReq.Bio = utils.Ptr(bio)
	}
	formFile, err := ctx.FormFile("picture")
	if err != nil {
		if err != http.ErrMissingFile {
			return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
		}
	}
	if formFile != nil {
		imageBytes, err := utils.ReadFormFile(formFile)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
		}
		updateReq.Picture = &share.File{
			File:     imageBytes,
			Filename: formFile.Filename,
		}
	}
	resp, err := s.profile.Update(context.TODO(), &updateReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *profileHandler) GetProfile(ctx echo.Context) error {
	var getReq profile.GetRequest
	getReq.UserId = ctx.Param("userId")
	p, err := s.profile.Get(context.TODO(), &getReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, p)
}

func (s *profileHandler) AddToList(ctx echo.Context) error {
	auth, ok := ctx.Get("claims").(*user.UserResponse)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, struct{ Message string }{"no credentials provided"})
	}
	if auth.UserId != ctx.Param("userId") {
		return ctx.JSON(http.StatusUnauthorized, struct{ Message string }{
			fmt.Sprintf("signed is as %s, set %s", auth.UserId, ctx.Param("userId")),
		})
	}

	var addToListReq profile.AddToListRequest
	addToListReq.ProfileId = ctx.Param("userId")

	if err := ctx.Bind(&addToListReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	res, err := s.profile.AddToList(context.TODO(), &addToListReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *profileHandler) RemoveFromList(ctx echo.Context) error {
	var removeFromListReq profile.RemoveFromListRequest
	removeFromListReq.ProfileId = ctx.Param("userId")
	if err := ctx.Bind(&removeFromListReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	res, err := s.profile.RemoveFromList(context.TODO(), &removeFromListReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *profileHandler) GetList(ctx echo.Context) error {
	var getListReq profile.GetListRequest
	getListReq.ProfileId = ctx.Param("userId")
	previewList, err := s.profile.GetList(context.TODO(), &getListReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	resp := make(map[uint32][]*profile.ListResponse_ListEntry)
	for k, v := range previewList.Lists {
		resp[k] = v.Value
	}
	return ctx.JSON(http.StatusOK, resp)
}
