package services

import (
	"github.com/flucas97/bookstore/users-api/model/users"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
)

type UserService struct{}

func (service *UserService) Create(user users.User) (*users.User, *errors_utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) Find(id int64) (*users.User, *errors_utils.RestErr) {
	result := &users.User{ID: id}
	if err := result.Find(); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUser patch user
func (service *UserService) Update(user users.User) (*users.User, *errors_utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser destoy a user
func (service *UserService) Delete(user *users.User) *errors_utils.RestErr {
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func (service UserService) Search(s string) (users.Users, *errors_utils.RestErr) {
	result, err := users.Search(s)
	if err != nil {
		return nil, err
	}

	return result, nil
}
