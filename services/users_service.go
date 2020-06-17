package services

import (
	"github.com/flucas97/bookstore/users-api/domain/users"
	"github.com/flucas97/bookstore/users-api/utils"
)

func CreateUser(user users.User) (*users.User, *utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUser(id int64) (*users.User, *utils.RestErr) {
	result := &users.User{ID: id}
	if err := result.Find(); err != nil {
		return nil, err
	}

	return result, nil
}
