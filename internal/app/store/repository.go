package store

import "restApi.com/users/viktorkorin/developer/testApps/testGo/restApi/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	ShowAll() ([]model.User, error)
}
