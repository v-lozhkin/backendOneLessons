package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

func (d delivery) Upload(ectx echo.Context) error {
	timer := time.Now()
	defer func() {
		d.stat.MethodDuration.With(prometheus.Labels{
			"method_name": "Upload",
		}).Observe(time.Since(timer).Seconds())
	}()

	ctx := ectx.Request().Context()
	filter := ItemFilter{}
	if err := ectx.Bind(&filter); err != nil {
		return err
	}

	existing, err := d.items.List(ctx, models.ItemFilter(filter))
	if err != nil {
		return convertToEchoError(err)
	}

	if len(existing) == 0 {
		return echo.ErrNotFound
	}

	fileHeader, err := ectx.FormFile("file")
	if err != nil {
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	path, err := d.images.Save(ctx, fileHeader.Filename, data)
	if err != nil {
		return convertToEchoError(err)
	}

	existing[0].ImageLink = path

	if err = d.items.Update(ctx, existing[0]); err != nil {
		return convertToEchoError(err)
	}

	extension := "undefined"
	if ext := filepath.Ext(fileHeader.Filename); ext != "" {
		extension = ext
	}

	d.stat.ExtensionCounter.With(prometheus.Labels{"extension": extension}).Inc()

	return ectx.JSON(http.StatusOK, existing[0])
}
