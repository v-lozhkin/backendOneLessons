package item

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrItemNotFound = errors.New("item not found")
)

type RESTDelivery interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type EchoDelivery interface {
	Create(ectx echo.Context) error
	List(ectx echo.Context) error
	Update(ectx echo.Context) error
	Delete(ectx echo.Context) error
	Upload(ectx echo.Context) error
}

type Usecase interface {
	Create(ctx context.Context, item *models.Item) error
	List(ctx context.Context, filter models.ItemFilter) (models.ItemList, error)
	Update(ctx context.Context, item models.Item) error
	Delete(ctx context.Context, id int64) error
}

type Repository interface {
	Create(ctx context.Context, item *models.Item) error
	List(ctx context.Context, filter models.ItemFilter) (models.ItemList, error)
	Update(ctx context.Context, item models.Item) error
	Delete(ctx context.Context, id int64) error
}
