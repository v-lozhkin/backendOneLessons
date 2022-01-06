package inmemory

import (
	itempkg "backendOneLessons/lesson4/internal/pkg/item"
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
)

type inmemory struct {
	iterator int // race unsafe
	items    []models.Item
}

func (in *inmemory) Create(_ context.Context, item *models.Item) error {
	item.ID = in.iterator
	in.iterator++

	in.items = append(in.items, *item)

	return nil
}

func (in inmemory) List(_ context.Context, filter models.ItemFilter) ([]models.Item, error) {
	res := make([]models.Item, 0, in.iterator)
	for _, val := range in.items {
		if filter.ID != 0 && filter.ID != val.ID {
			continue
		}

		if filter.PriceMax != 0 && filter.PriceMax < val.Price {
			continue
		}

		if filter.PriceMin != 0 && filter.PriceMin > val.Price {
			continue
		}

		res = append(res, val)
	}

	return res, nil
}

func (in *inmemory) Update(_ context.Context, item models.Item) error {
	found := false
	for i, itm := range in.items {
		if itm.ID == item.ID {
			in.items = append(append(in.items[:i], item), in.items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return itempkg.ErrItemNotFound
	}

	return nil
}

func (in *inmemory) Delete(ctx context.Context, id int) error {
	found := false
	for i, itm := range in.items {
		if itm.ID == id {
			in.items = append(in.items[:i], in.items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return itempkg.ErrItemNotFound
	}

	return nil
}

func New() itempkg.ItemUsecase {
	return &inmemory{
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
