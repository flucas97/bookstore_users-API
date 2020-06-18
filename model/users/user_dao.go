package users

import (
	"fmt"
	"strings"

	"github.com/flucas97/bookstore/users-api/datasources/mysql/users_db"
	"github.com/flucas97/bookstore/users-api/utils"
)

const (
	queryInsertUser  = ("INSERT INTO users(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?);")
	queryFindUser    = ("SELECT * FROM users WHERE id=?;")
	indexUniqueEmail = "email_UNIQUE"
)

var (
	usersDB = make(map[int64]*User)
)

// Save persist user in database
func (user *User) Save() *utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Error: %v", err))
	}
	defer stmt.Close() // Close db connection with this statement

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return utils.NewBadRequestError(fmt.Sprintf("user: '%v' already registered", user.Email))
		}
		return utils.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.CreatedAt = utils.GetNowString()
	user.ID = userID
	return nil
}

// Find gets a user
func (user *User) Find() *utils.RestErr {
	// check if is everything OK accessing the DB
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
