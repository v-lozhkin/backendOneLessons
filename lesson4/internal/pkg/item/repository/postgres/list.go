package postgres

import (
	repomodels "backendOneLessons/lesson4/internal/pkg/item/repository/models"
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"fmt"
	"strings"
)

func (r repository) List(ctx context.Context, filter models.ItemFilter) (models.ItemList, error) {
	defer r.stat.MethodDuration.WithLabels(map[string]string{"method_name": "List"}).Start().Stop()

	res := repomodels.ItemList{}

	query := strings.Builder{}
	query.WriteString("SELECT * FROM item where true")
	args := make([]interface{}, 0)

	if filter.ID != nil {
		query.WriteString(" and id = ?")
		args = append(args, *filter.ID)
	}
	if filter.PriceMax != nil {
		query.WriteString(" and price <=  ?")
		args = append(args, *filter.PriceMax)
	}
	if filter.PriceMin != nil {
		query.WriteString(" and price >=")
		args = append(args, *filter.PriceMin)
	}

	if err := r.db.SelectContext(ctx, &res, query.String(), args...); err != nil {
		return nil, fmt.Errorf("failed to select item from db: %w", err)
	}

	return repomodels.RepoItemListToModel(res), nil
}
