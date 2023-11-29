package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/profile"
	"net/http"
)

func (s *server) UpdateProfile(ctx echo.Context) error {
	var updateReq profile.UpdateProfileRequest
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
		updateReq.Picture = imageBytes
	}
	resp, err := s.profile.UpdateProfile(context.TODO(), &updateReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}
