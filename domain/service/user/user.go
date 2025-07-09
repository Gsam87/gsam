package userservice

import (
	"errors"
)

type UserService interface {
	CheckID(id string) error
}

type UserServiceImpl struct {
}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) CheckID(
	id string,
) error {
	if len(id) > 0 {
		return nil
	} else {
		return errors.New("user not found")
	}
}
