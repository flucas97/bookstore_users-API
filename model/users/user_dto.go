package users

import (
	"strings"

	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

func (u *User) Validate() *errors_utils.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return errors_utils.NewBadRequestError("Invalid email address")
	}

	return nil
}
