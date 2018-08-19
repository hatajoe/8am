package v1

import (
	"context"

	"github.com/hatajoe/8am/app/interface/rpc/v1.0/protocol"
	"github.com/hatajoe/8am/app/usecase"
)

type userService struct {
	userUsecase usecase.UserUsecase
}

func NewUserService(userUsecase usecase.UserUsecase) *userService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (s *userService) ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error) {
	users, err := s.userUsecase.ListUser()
	if err != nil {
		return nil, err
	}

	res := &protocol.ListUserResponseType{
		Users: toUser(users),
	}

	return res, nil
}

func (s *userService) RegisterUser(ctx context.Context, in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error) {
	if err := s.userUsecase.RegisterUser(in.GetEmail()); err != nil {
		return &protocol.RegisterUserResponseType{}, err
	}
	return &protocol.RegisterUserResponseType{}, nil
}

func toUser(users []*usecase.User) []*protocol.User {
	res := make([]*protocol.User, len(users))
	for i, user := range users {
		res[i] = &protocol.User{
			Id:    user.ID,
			Email: user.Email,
		}
	}
	return res
}
