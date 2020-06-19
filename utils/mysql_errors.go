package utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *RestErr {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return NewNotFoundError("no record matching given id")
		}
		return NewInternalServerError(fmt.Sprintf("error parsing database response"))
	}

	switch mysqlErr.Number {
	case 1062:
		return NewBadRequestError("email already exists")
	}

	return NewInternalServerError(fmt.Sprintf("error processing request"))
}
