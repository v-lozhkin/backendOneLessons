package postgres

import (
	repomodels "backendOneLessons/lesson4/internal/pkg/item/repository/models"
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"fmt"
)

func (r repository) Create(ctx context.Context, item *models.Item) error {
	defer r.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Create"}).Start().Stop()

	query := "INSERT INTO item (name, description, price, image_link) VALUES " +
		"(:name, :description, :price, :image_link) RETURNING id"
	res, err := r.db.NamedQueryContext(ctx,
		query,
		repomodels.ModelToRepoItem(*item),
	)
	if err != nil {
		return fmt.Errorf("failed to insert item to db: %w", err)
	}

	var id int64

	next := res.Next()
	if err = res.Scan(&id); err != nil || !next {
		return fmt.Errorf("failed to get last inserted id from db: %w (next: %t)", err, next)
	}

	item.ID = id
	return nil
}
