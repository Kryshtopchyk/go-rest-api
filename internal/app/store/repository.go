package store

import "github.com/Kryshtopchyk/go-rest-api/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
