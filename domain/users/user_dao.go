package users

import (
	"fmt"

	"github.com/flucas97/bookstore/users-api/utils"
)

var (
	usersDB = make(map[int64]*User)
)

// Save persist user in database
func (user User) Save() *utils.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return utils.NewBadRequestError(fmt.Sprintf("User %v already registered", user.Email))
		}
		return utils.NewBadRequestError(fmt.Sprintf("User %v already exists", user.ID))
	}

	usersDB[user.ID] = &user
	return nil
}

// Find gets a user
func (user *User) Find() *utils.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return utils.NewNotFoundError(fmt.Sprintf("User %v not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.Email = result.Email
	user.LastName = result.LastName
	user.CreatedAt = result.CreatedAt

	return nil
}
