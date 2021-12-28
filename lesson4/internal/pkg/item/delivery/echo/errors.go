package echo

import (
	"backendOneLessons/lesson4/internal/pkg/item"
	"errors"

	"github.com/labstack/echo/v4"
)

func convertToEchoError(err error) error {
	if errors.Is(err, item.ErrItemNotFound) {
		return echo.ErrNotFound
	}

	return err
}
