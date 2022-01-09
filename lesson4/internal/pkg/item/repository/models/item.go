package models

import "database/sql"

type Item struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	Price       int            `db:"price"`
	ImageLink   sql.NullString `db:"image_link"`
}

type ItemList []Item
