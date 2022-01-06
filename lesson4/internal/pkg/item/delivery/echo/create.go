package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

func (d delivery) Create(ectx echo.Context) error {
	timer := time.Now()
	defer func() {
		d.stat.MethodDuration.With(prometheus.Labels{
			"method_name": "Create",
		}).Observe(time.Since(timer).Seconds())
	}()

	newItem := &Item{}
	if err := ectx.Bind(newItem); err != nil {
		return err
	}

	if err := d.items.Create(ectx.Request().Context(), (*models.Item)(newItem)); err != nil {
		return convertToEchoError(err)
	}

	return ectx.JSON(http.StatusOK, newItem)
}
