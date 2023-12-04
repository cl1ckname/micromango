package middleware

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/user"
	"net/http"
	"regexp"
)

var bearerRegexp = regexp.MustCompile(`Bearer\s([\w-]*\.[\w-]*\.[\w-]*$)`)

func Auth(u user.UserClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			log.Println("header", header)
			if header != "" {
				token, ok := parseBearer(header)
				if !ok {
					return utils.ErrorToResponse(c, fmt.Errorf("invalid token"))
				}
				claims, err := u.Auth(context.TODO(), &user.AuthRequest{Token: token})
				if err != nil {
					return c.String(http.StatusUnauthorized, err.Error())
				}
				c.Set("claims", claims)
			}
			c.Set("a", "b")
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
