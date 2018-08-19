package service

import (
	"fmt"

	"github.com/hatajoe/8am/app/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Duplicated(email string) error {
	user, err := s.repo.FindByEmail(email)
	if user != nil {
		return fmt.Errorf("%s already exists", email)
	}
	if err != nil {
		return err
	}
	return nil
}
