package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

func (d delivery) List(ectx echo.Context) error {
	timer := time.Now()
	defer func() {
		d.stat.MethodDuration.With(prometheus.Labels{
			"method_name": "List",
		}).Observe(time.Since(timer).Seconds())
	}()

	time.Sleep(time.Millisecond * 490)

	filter := ItemFilter{}
	if err := ectx.Bind(&filter); err != nil {
		return err
	}

	list, err := d.items.List(ectx.Request().Context(), models.ItemFilter(filter))
	if err != nil {
		return convertToEchoError(err)
	}
	if len(list) == 0 && filter.ID != 0 {
		return echo.ErrNotFound
	}

	if filter.ID == 0 {
		return ectx.JSON(http.StatusOK, list)
	}

	return ectx.JSON(http.StatusOK, list[0])
}