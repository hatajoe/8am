package memory

import (
	"sync"

	"github.com/hatajoe/8am/app/domain/model"
)

type userRepository struct {
	mu    *sync.Mutex
	users map[string]*User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		mu:    &sync.Mutex{},
		users: map[string]*User{},
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]*model.User, len(r.users))
	i := 0
	for _, user := range r.users {
		users[i] = model.NewUser(user.ID, user.Email)
		i++
	}
	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Email == email {
			return model.NewUser(user.ID, user.Email), nil
		}
	}
	return nil, nil
}

func (r *userRepository) Save(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.GetID()] = &User{
		ID:    user.GetID(),
		Email: user.GetEmail(),
	}
	return nil
}

type User struct {
	ID    string
	Email string
}
