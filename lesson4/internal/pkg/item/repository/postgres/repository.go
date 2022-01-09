package postgres

import (
	"backendOneLessons/lesson4/internal/pkg/item"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) item.Repository {
	return repository{db: db}
}
