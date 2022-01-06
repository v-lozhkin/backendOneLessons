package usecase

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"fmt"
)

func (u usecase) List(ctx context.Context, filter models.ItemFilter) ([]models.Item, error) {
	if err := filter.Validate(); err != nil {
		return nil, fmt.Errorf("item's validate failed: %w", err)
	}

	list, err := u.repo.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list items from repo: %w", err)
	}

	return list, nil
}
