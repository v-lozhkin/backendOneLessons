package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d delivery) List(ectx echo.Context) error {
	defer d.stat.MethodDuration.WithLabels(map[string]string{"method_name": "List"}).Start().Stop()

	filter := ItemFilter{}
	if err := ectx.Bind(&filter); err != nil {
		return err
	}

	list, err := d.items.List(ectx.Request().Context(), models.ItemFilter(filter))
	if err != nil {
		return convertToEchoError(err)
	}
	if len(list) == 0 && filter.ID != nil && *filter.ID != 0 {
		return echo.ErrNotFound
	}

	if filter.ID == nil || *filter.ID == 0 {
		return ectx.JSON(http.StatusOK, list)
	}

	return ectx.JSON(http.StatusOK, list[0])
}
