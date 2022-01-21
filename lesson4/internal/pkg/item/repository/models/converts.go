package models

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"database/sql"
)

func ModelToRepoItem(item models.Item) Item {
	return Item{
		ID:   item.ID,
		Name: item.Name,
		Description: sql.NullString{
			String: item.Description,
			Valid:  true,
		},
		Price: item.Price,
		ImageLink: sql.NullString{
			String: item.ImageLink,
			Valid:  true,
		},
	}
}

func RepoItemToModel(item Item) models.Item {
	return models.Item{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description.String,
		Price:       item.Price,
		ImageLink:   item.ImageLink.String,
	}
}

func RepoItemListToModel(items ItemList) models.ItemList {
	res := make(models.ItemList, 0, len(items))

	for _, itm := range items {
		res = append(res, RepoItemToModel(itm))
	}

	return res
}
