package usecase

import "backendOneLessons/lesson4/internal/pkg/user"

type inmemory struct {
	users map[string]string
}

func (i inmemory) Validate(login, password string) bool {
	for name, pass := range i.users {
		if name == login && pass == password {
			return true
		}
	}

	return false
}

func New(users map[string]string) user.Usecase {
	return inmemory{users: users}
}
