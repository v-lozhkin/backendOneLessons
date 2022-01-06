package middlewares

import (
	"backendOneLessons/lesson4/internal/pkg/user"
	"backendOneLessons/lesson4/internal/pkg/user/delivery"
	"encoding/base64"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuthMiddlewareFull(users user.Usecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			auth64 := strings.TrimPrefix(ectx.Request().Header.Get(echo.HeaderAuthorization), "Basic ")

			//admin:password
			auth, err := base64.StdEncoding.DecodeString(auth64)
			if err != nil {
				return echo.ErrUnauthorized
			}

			credentials := strings.Split(string(auth), ":")
			if len(credentials) != 2 {
				return echo.ErrUnauthorized
			}

			username, pass := credentials[0], credentials[1]

			if !users.Validate(username, pass) {
				return echo.ErrUnauthorized
			}

			return next(ectx)
		}
	}
}

func BasicAuthMiddleware(users user.Usecase) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, ectx echo.Context) (bool, error) {
		if !users.Validate(username, password) {
			return false, echo.ErrUnauthorized
		}

		return true, nil
	})
}

func JWTAuthMiddleware(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: nil,
		SigningKey:     []byte(secret),
		Claims:         &delivery.Payload{},
	})
}
