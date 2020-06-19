package users

import (
	"fmt"
	"strings"

	"github.com/flucas97/bookstore/users-api/datasources/mysql/users_db"
	"github.com/flucas97/bookstore/users-api/utils"
)

const (
	queryUpdateUser = ("UPDATE users SET first_name=?, last_name=?, email=?, created_at=?, updated_at=? WHERE id=?;")
	queryInsertUser = ("INSERT INTO users(first_name, last_name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?);")
	queryFindUser   = ("SELECT id, first_name, last_name, email, created_at, updated_at FROM users WHERE id=?;")
)

// Save persist user in database
func (user *User) Save() *utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Error: %v", err))
	}
	defer stmt.Close() // Close db connection with this statement

	user.CreatedAt, user.UpdatedAt = utils.GetNowString(), utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return utils.ParseError(err)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.ID = userID
	return nil
}

// Find gets a user
func (user *User) Find() *utils.RestErr {
	// check if is everything OK accessing the DB
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryFindUser)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintln("error while preparing search query"))
	}
	defer stmt.Close()

	searchResult := stmt.QueryRow(user.ID)

	if err := searchResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return utils.NewNotFoundError(fmt.Sprintf("user ID:%v not found", user.ID))
		}
		return utils.NewInternalServerError(fmt.Sprintf("error trying to get user %v error: %s", user.ID, err.Error()))
	}

	return nil
}

// Update a existent user
func (user *User) Update() *utils.RestErr {
	// montar a query
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintln("error while preparing search query"))
	}
	defer stmt.Close()

	user.UpdatedAt = utils.GetNowString()
	// executa
	_, err = stmt.Exec(&user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.ID)
	if err != nil {
		return utils.ParseError(err)
	}
	fmt.Println(">>>>>>>", user)
	return nil
}
