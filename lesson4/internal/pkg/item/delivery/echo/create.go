package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d delivery) Create(ectx echo.Context) error {
	defer d.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Create"}).Start().Stop()

	newItem := &Item{}
	if err := ectx.Bind(newItem); err != nil {
		return err
	}

	if err := d.items.Create(ectx.Request().Context(), (*models.Item)(newItem)); err != nil {
		return convertToEchoError(err)
	}

	return ectx.JSON(http.StatusOK, newItem)
}
