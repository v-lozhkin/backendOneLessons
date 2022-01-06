package usecase

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"fmt"
)

func (u usecase) Create(ctx context.Context, item *models.Item) error {
	if err := item.Validate(); err != nil {
		return fmt.Errorf("item's validate failed: %w", err)
	}

	if err := u.repo.Create(ctx, item); err != nil {
		return fmt.Errorf("failed to create item in repo: %w", err)
	}

	return nil
}
