package echo

import (
	"backendOneLessons/lesson4/internal/pkg/image"
	"backendOneLessons/lesson4/internal/pkg/item"
	"backendOneLessons/lesson4/internal/pkg/models"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	items  item.ItemUsecase
	images image.Storage
}

func New(items item.ItemUsecase, images image.Storage) item.EchoDelivery {
	return delivery{
		items:  items,
		images: images,
	}
}

func (d delivery) Create(ectx echo.Context) error {
	newItem := &Item{}
	if err := ectx.Bind(newItem); err != nil {
		return err
	}

	if err := d.items.Create(ectx.Request().Context(), (*models.Item)(newItem)); err != nil {
		return convertToEchoError(err)
	}

	return ectx.JSON(http.StatusOK, newItem)
}

func (d delivery) List(ectx echo.Context) error {
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

func (d delivery) Update(ectx echo.Context) error {
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

func (d delivery) Delete(ectx echo.Context) error {
	filter := ItemFilter{}
	if err := ectx.Bind(&filter); err != nil {
		return err
	}

	return convertToEchoError(d.items.Delete(ectx.Request().Context(), filter.ID))
}

func (d delivery) Upload(ectx echo.Context) error {
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

	return ectx.JSON(http.StatusOK, existing[0])
}
