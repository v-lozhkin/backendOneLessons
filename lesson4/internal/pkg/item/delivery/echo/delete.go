package echo

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

func (d delivery) Delete(ectx echo.Context) error {
	timer := time.Now()
	defer func() {
		d.stat.MethodDuration.With(prometheus.Labels{
			"method_name": "Delete",
		}).Observe(time.Since(timer).Seconds())
	}()

	filter := ItemFilter{}
	if err := ectx.Bind(&filter); err != nil {
		return err
	}

	return convertToEchoError(d.items.Delete(ectx.Request().Context(), filter.ID))
}
