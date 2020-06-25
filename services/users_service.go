package services

import (
	"github.com/flucas97/bookstore/users-api/model/users"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
)

var (
	UsersService UsersServiceInterface = &usersService{}
)

type usersService struct{}

type UsersServiceInterface interface {
	Create(users.User) (*users.User, *errors_utils.RestErr)
	Find(int64) (*users.User, *errors_utils.RestErr)
	Update(users.User) (*users.User, *errors_utils.RestErr)
	Delete(*users.User) *errors_utils.RestErr
	Search(string) (users.Users, *errors_utils.RestErr)
	LoginUser() (*users.User, *errors_utils.RestErr)
}

func (service *usersService) Create(user users.User) (*users.User, *errors_utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *usersService) Find(id int64) (*users.User, *errors_utils.RestErr) {
	result := &users.User{ID: id}
	if err := result.Find(); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUser patch user
func (service *usersService) Update(user users.User) (*users.User, *errors_utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser destoy a user
func (service *usersService) Delete(user *users.User) *errors_utils.RestErr {
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func (service usersService) Search(s string) (users.Users, *errors_utils.RestErr) {
	result, err := users.Search(s)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service usersService) LoginUser() (*users.User, *errors_utils.RestErr) {
	return nil, nil
}
