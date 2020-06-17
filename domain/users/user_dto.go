package users

import (
	"strings"

	"github.com/flucas97/bookstore/users-api/utils"
)

type User struct {
	ID        int64  `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (u *User) Validate() *utils.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return utils.NewBadRequestError("Invalid email address")
	}

	return nil
}
