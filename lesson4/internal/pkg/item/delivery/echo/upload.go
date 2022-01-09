package echo

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func (d delivery) Upload(ectx echo.Context) error {
	defer d.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Upload"}).Start().Stop()

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

	d.stat.ExtensionCounter.With(map[string]string{"extension": extension}).Inc()

	return ectx.JSON(http.StatusOK, existing[0])
}
