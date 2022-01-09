package delivery

import (
	"backendOneLessons/lesson4/internal/pkg/user"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type delivery struct {
	users     user.Usecase
	ttl       int64
	secretKey []byte
}

func New(users user.Usecase, ttl int64, secret string) user.Delivery {
	return delivery{
		users:     users,
		ttl:       ttl,
		secretKey: []byte(secret),
	}
}

type Payload struct {
	jwt.StandardClaims

	Name string
}

type Credentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (d delivery) Login(ectx echo.Context) error {
	creds := Credentials{}
	if err := ectx.Bind(&creds); err != nil {
		return echo.ErrUnauthorized
	}

	if !d.users.Validate(creds.Name, creds.Password) {
		return echo.ErrUnauthorized
	}

	payload := Payload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: d.ttl,
		},
		Name: creds.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &payload)

	signedToken, err := token.SignedString(d.secretKey)
	if err != nil {
		return err
	}

	ectx.Response().Header().Set("X-Expires-After", time.Unix(d.ttl, 0).String())

	return ectx.JSON(http.StatusOK, signedToken)
}
