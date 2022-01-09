package usecase

import (
	"context"
	"fmt"
)

func (u usecase) Delete(ctx context.Context, id int64) error {
	defer u.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Delete"}).Start().Stop()
	if err := u.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete item in repo: %w", err)
	}

	return nil
}
