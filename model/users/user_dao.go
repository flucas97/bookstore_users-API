package users

import (
	"fmt"
	"strings"

	"github.com/flucas97/bookstore/users-api/datasources/mysql/users_db"
	"github.com/flucas97/bookstore/users-api/logger"
	"github.com/flucas97/bookstore/users-api/utils/crypto_utils"
	"github.com/flucas97/bookstore/users-api/utils/dates_utils"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
	"github.com/flucas97/bookstore/users-api/utils/mysql_utils"
)

const (
	queryUpdateUser             = ("UPDATE users SET first_name=?, last_name=?, email=?, password=?, created_at=?, updated_at=? WHERE id=?;")
	queryInsertUser             = ("INSERT INTO users(first_name, last_name, email, password, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?);")
	queryFindUser               = ("SELECT id, first_name, last_name, email, status, created_at, updated_at FROM users WHERE id=?;")
	queryDeleteUser             = ("DELETE FROM users WHERE id=?;")
	queryFindByStatus           = ("SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users WHERE status=?;")
	queryFindByEmailAndPassword = ("SELECT id, first_name, last_name, email, created_at, updated_at, status FROM users WHERE email=? AND password=?;")
	statusActive                = "active"
	statusEnded                 = "ended"
)

// Save persist user in database
func (user *User) Save() *errors_utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error while preparing Save query", err)
		return errors_utils.NewInternalServerError("an error occurred while saving user. Try again")
	}
	defer stmt.Close() // Close db connection with this statement

	user.CreatedAt, user.UpdatedAt, user.Status = dates_utils.GetNowString(), dates_utils.GetNowString(), statusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors_utils.NewBadRequestError(fmt.Sprintf("user '%v' already exists", user.Email))
		}
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error while saving user in database query", err)
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
		logger.Error("error while preparing Find query", err)
		return errors_utils.NewInternalServerError("an error occurred while finding user. Try again")
	}
	defer stmt.Close()

	searchResult := stmt.QueryRow(user.ID)

	if err := searchResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors_utils.NewNotFoundError(fmt.Sprintf("user ID:%v not found", user.ID))
		}
		logger.Error("error trying to find user in database", err)
		return errors_utils.NewInternalServerError(fmt.Sprintf("error trying to get user %v", user.ID))
	}
	return nil
}

// Update a existent user
func (user *User) Update() *errors_utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error while preparing search query", err)
		return errors_utils.NewInternalServerError("an error occurred while updating user. Try again")
	}
	defer stmt.Close()

	user.UpdatedAt = dates_utils.GetNowString()

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
		logger.Error("error while preparing Delete query", err)
		return errors_utils.NewInternalServerError("an error occurred while deleting user. Try again")
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

	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error while preparing Search query", err)
		return nil, errors_utils.NewInternalServerError("an error occurred while searching users. Try again")
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

// Find gets a user
func (user *User) FindUserByEmailAndPassword() *errors_utils.RestErr {
	// check if is everything OK accessing the DB
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error while preparing Find by Email and Password query", err)
		return errors_utils.NewInternalServerError("database error")
	}
	defer stmt.Close()

	searchResult := stmt.QueryRow(user.Email, user.Password)

	if err := searchResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Status); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors_utils.NewNotFoundError("invalid email address or password")
		}
		logger.Error("error trying to get user by email and password %v", err)
		return errors_utils.NewInternalServerError("database error")
	}
	return nil
}
