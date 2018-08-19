package repository

import "github.com/hatajoe/8am/app/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(*model.User) error
}
