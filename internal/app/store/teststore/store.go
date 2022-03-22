package teststore

import (
	"restApi.com/users/viktorkorin/developer/testApps/testGo/restApi/internal/app/model"
	"restApi.com/users/viktorkorin/developer/testApps/testGo/restApi/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
