package users

import (
	"fmt"
	"strings"

	"github.com/flucas97/bookstore/users-api/datasources/mysql/users_db"
	"github.com/flucas97/bookstore/users-api/utils/dates_utils"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
	"github.com/flucas97/bookstore/users-api/utils/mysql_utils"
)

const (
	queryUpdateUser        = ("UPDATE users SET first_name=?, last_name=?, email=?, password=?, created_at=?, updated_at=? WHERE id=?;")
	queryInsertUser        = ("INSERT INTO users(first_name, last_name, email, password, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?);")
	queryFindUser          = ("SELECT id, first_name, last_name, email, status, created_at, updated_at FROM users WHERE id=?;")
	queryDeleteUser        = ("DELETE FROM users WHERE id=?;")
	queryFindUsersByStatus = ("SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users WHERE status=?;")
	statusActive           = "active"
	statusEnded            = "ended"
)

// Save persist user in database
func (user *User) Save() *errors_utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors_utils.NewInternalServerError(fmt.Sprintf("Error: %v", err))
	}
	defer stmt.Close() // Close db connection with this statement

	user.CreatedAt, user.UpdatedAt, user.Status = dates_utils.GetNowString(), dates_utils.GetNowString(), statusActive

	insertResult, _ := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors_utils.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.ID = userID
	return nil
}

// Find gets a user
func (user *User) Find() *errors_utils.RestErr {
	// check if is everything OK accessing the DB
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryFindUser)
	if err != nil {
		return errors_utils.NewInternalServerError(fmt.Sprintln("error while preparing search query"))
	}
	defer stmt.Close()

	searchResult := stmt.QueryRow(user.ID)

	if err := searchResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors_utils.NewNotFoundError(fmt.Sprintf("user ID:%v not found", user.ID))
		}
		return errors_utils.NewInternalServerError(fmt.Sprintf("error trying to get user %v error: %s", user.ID, err.Error()))
	}
	return nil
}

// Update a existent user
func (user *User) Update() *errors_utils.RestErr {
	// montar a query
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors_utils.NewInternalServerError(fmt.Sprintln("error while preparing search query"))
	}
	defer stmt.Close()

	user.UpdatedAt = dates_utils.GetNowString()

	// executa
	_, err = stmt.Exec(&user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

// Delete destroy a user
func (user *User) Delete() *errors_utils.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors_utils.NewInternalServerError(fmt.Sprintln("error while preparing search query"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(&user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

// FindByStatus get all active users
func Search(status string) ([]User, *errors_utils.RestErr) {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryFindUsersByStatus)
	if err != nil {
		return nil, errors_utils.NewInternalServerError(fmt.Sprintln("error while preparing search query"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	usersResult := make([]User, 0)

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status)
		if err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		usersResult = append(usersResult, user)
	}

	if len(usersResult) == 0 {
		return nil, errors_utils.NewNotFoundError(fmt.Sprintf("no users with status '%v' found", status))
	}

	return usersResult, nil
}
