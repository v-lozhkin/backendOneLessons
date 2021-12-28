package user

import "github.com/labstack/echo/v4"

type Delivery interface {
	Login(ectx echo.Context) error
}

type Usecase interface {
	Validate(login, password string) bool
}
