package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d delivery) Update(ectx echo.Context) error {
	defer d.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Update"}).Start().Stop()

	request := struct {
		Item
		ItemFilter
	}{}
	if err := ectx.Bind(&request); err != nil {
		return err
	}

	if request.ItemFilter.ID == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "item id can't be empty")
	}

	request.Item.ID = *request.ItemFilter.ID

	return d.items.Update(ectx.Request().Context(), models.Item(request.Item))
}
