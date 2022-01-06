package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

func (d delivery) Update(ectx echo.Context) error {
	timer := time.Now()
	defer func() {
		d.stat.MethodDuration.With(prometheus.Labels{
			"method_name": "Update",
		}).Observe(time.Since(timer).Seconds())
	}()

	request := struct {
		Item
		ItemFilter
	}{}
	if err := ectx.Bind(&request); err != nil {
		return err
	}

	request.Item.ID = request.ItemFilter.ID

	return d.items.Update(ectx.Request().Context(), models.Item(request.Item))
}
