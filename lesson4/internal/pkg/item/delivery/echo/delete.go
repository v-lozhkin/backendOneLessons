package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d delivery) Delete(ectx echo.Context) error {
	defer d.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Delete"}).Start().Stop()

	filter := ItemFilter{}
	if err := ectx.Bind(&filter); err != nil {
		return err
	}

	if filter.ID == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "item id can't be empty")
	}

	return convertToEchoError(d.items.Delete(ectx.Request().Context(), *filter.ID))
}
