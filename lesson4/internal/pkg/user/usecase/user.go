package usecase

import "backendOneLessons/lesson4/internal/pkg/user"

type inmemory struct {
	repo user.Repository
}

func (i inmemory) Validate(login, password string) bool {
	for _, usr := range i.repo.List() {
		if usr.Login == login && usr.Password == password {
			return true
		}
	}

	return false
}

func New(repo user.Repository) user.Usecase {
	return inmemory{repo: repo}
}
