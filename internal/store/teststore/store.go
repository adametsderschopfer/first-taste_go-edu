package teststore

import (
	"github.com/adametsderschopfer/http-rest-api_go-edu/internal/store"
	"github.com/adametsderschopfer/http-rest-api_go-edu/model"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return nil
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
