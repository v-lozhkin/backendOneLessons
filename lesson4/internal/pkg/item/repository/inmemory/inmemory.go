package inmemory

import (
	itempkg "backendOneLessons/lesson4/internal/pkg/item"
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
)

type repository struct {
	iterator int64 // race unsafe
	items    []models.Item
}

func (r *repository) Create(_ context.Context, item *models.Item) error {
	item.ID = r.iterator
	r.iterator++

	r.items = append(r.items, *item)

	return nil
}

func (r repository) List(_ context.Context, filter models.ItemFilter) (models.ItemList, error) {
	res := make([]models.Item, 0, r.iterator)
	for _, val := range r.items {
		if filter.ID != nil && *filter.ID != val.ID {
			continue
		}

		if filter.PriceMax != nil && *filter.PriceMax < val.Price {
			continue
		}

		if filter.PriceMin != nil && *filter.PriceMin > val.Price {
			continue
		}

		res = append(res, val)
	}

	return res, nil
}

func (r *repository) Update(_ context.Context, item models.Item) error {
	found := false
	for i, itm := range r.items {
		if itm.ID == item.ID {
			r.items = append(append(r.items[:i], item), r.items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return itempkg.ErrItemNotFound
	}

	return nil
}

func (r *repository) Delete(_ context.Context, id int64) error {
	found := false
	for i, itm := range r.items {
		if itm.ID == id {
			r.items = append(r.items[:i], r.items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return itempkg.ErrItemNotFound
	}

	return nil
}

func New() itempkg.Usecase {
	return &repository{
		items: []models.Item{
			{
				ID:          1,
				Name:        "Snack",
				Description: "Just some snack",
				Price:       100,
			},
		},
		iterator: 2,
	}
}
