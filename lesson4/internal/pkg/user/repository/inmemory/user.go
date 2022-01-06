package inmemory

import (
	"backendOneLessons/lesson4/internal/pkg/models"
	"backendOneLessons/lesson4/internal/pkg/user"
)

type inmemory struct{}

func (i inmemory) List() []models.User {
	return []models.User{
		{
			Login:    "admin",
			Password: "password",
		},
	}
}

func New() user.Repository {
	return inmemory{}
}
