package middleware

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
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
				var ok bool
				token, ok = parseBearer(header)
				if !ok {
					return utils.ErrorToResponse(c, fmt.Errorf("invalid token"))
				}
			} else {
				cookie, err := c.Cookie("auth")
				if err != nil {
					return utils.ErrorToResponse(c, err)
				}
				token = cookie.Value
			}
			req := &user.AuthRequest{Token: token}
			claims, err := u.Auth(context.TODO(), req)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, struct{ Message string }{err.Error()})
			}
			c.Set("claims", claims)
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
