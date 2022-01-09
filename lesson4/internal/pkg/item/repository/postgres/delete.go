package postgres

import (
	"context"
	"fmt"
)

func (r repository) Delete(ctx context.Context, id int64) error {
	defer r.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Delete"}).Start().Stop()

	stmt, err := r.db.Prepare("DELETE FROM item WHERE id= $1")
	if err != nil {
		return fmt.Errorf("failed to prepare delete statement: %w", err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete %d item from repo: %w", id, err)
	}
	return nil
}
