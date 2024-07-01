package store

import "github.com/adametsderschopfer/http-rest-api_go-edu/model"

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
