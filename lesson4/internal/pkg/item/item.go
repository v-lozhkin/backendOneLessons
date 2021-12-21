package item

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"net/http"
)

type RESTDelivery interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type ItemUsecase interface {
	Create(ctx context.Context, item *models.Item) error
	List(ctx context.Context, filter models.ItemFilter) ([]models.Item, error)
	Update(ctx context.Context, item models.Item) error
	Delete(ctx context.Context, id int) error
}
