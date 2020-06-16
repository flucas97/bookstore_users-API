package services

import (
	"net/http"

	"github.com/flucas97/bookstore/users-api/domain/users"
	"github.com/flucas97/bookstore/users-api/utils"
)

func CreateUser(u users.User) (*users.User, *utils.RestErr) {
	return &u, &utils.RestErr{
		Status: http.StatusInternalServerError,
	}
}

func FindUser() {
}
