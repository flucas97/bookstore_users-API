package services

import (
	"github.com/flucas97/bookstore/users-api/domain/users"
)

func CreateUser(u users.User) (*users.User, error) {
	return &u, nil
}

func FindUser() {
}
