package usecase

import (
	"backendOneLessons/lesson4/internal/pkg/item"
)

type usecase struct {
	repo item.Repository
}

func New(repo item.Repository) item.Usecase {
	return usecase{repo: repo}
}
