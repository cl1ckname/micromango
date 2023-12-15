package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/user"
	"net/http"
	"regexp"
)

var bearerRegexp = regexp.MustCompile(`Bearer\s([\w-]*\.[\w-]*\.[\w-]*$)`)

func Auth(u user.UserClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var token string
			header := c.Request().Header.Get("Authorization")
			if header != "" {
				log.Printf("%s - header is not null\n", c.Request().URL.Path)
				var ok bool
				token, ok = parseBearer(header)
				if !ok {
					return utils.ErrorToResponse(c, fmt.Errorf("invalid token"))
				}
			} else {
				cookie, err := c.Cookie("auth")
				if err != nil {
					if !errors.Is(err, http.ErrNoCookie) {
						return utils.ErrorToResponse(c, fmt.Errorf("authentication error: %v", err))
					}
				} else {
					token = cookie.Value
				}
			}
			if token != "" {
				req := &user.AuthRequest{Token: token}
				claims, err := u.Auth(context.TODO(), req)
				if err != nil {
					return c.JSON(http.StatusUnauthorized, struct{ Message string }{err.Error()})
				}
				c.Set("claims", claims)
				log.Printf("%s - authorized as %s\n", c.Request().URL.Path, claims.UserId)
			}
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

func parseBearer(header string) (string, bool) {
	subexps := bearerRegexp.FindStringSubmatch(header)
	if len(subexps) != 2 {
		return "", false
	}
	return subexps[1], true
}
