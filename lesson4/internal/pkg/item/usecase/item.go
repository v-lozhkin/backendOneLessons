package usecase

import (
	"backendOneLessons/lesson4/internal/pkg/item"
)

type usecase struct {
	repo item.Repository
}

func New(repo item.Repository) item.ItemUsecase {
	return usecase{repo: repo}
}
