package postgres

import (
	repomodels "backendOneLessons/lesson4/internal/pkg/item/repository/models"
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"fmt"
)

// Update - accept just full state. Partial update not accepted
func (r repository) Update(ctx context.Context, item models.Item) error {
	query := "UPDATE item SET name=:name, description=:description, price=:price, image_link=:image_link WHERE ID=:id"

	_, err := r.db.NamedExecContext(ctx, query, repomodels.ModelToRepoItem(item))
	return fmt.Errorf("failed to update item %d: %w", item.ID, err)
}
