package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors_utils.RestErr {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors_utils.NewBadRequestError("no record matching given id")
		}
		return errors_utils.NewInternalServerError(fmt.Sprintf("error parsing database response"))
	}

	switch mysqlErr.Number {
	case 1062:
		return errors_utils.NewBadRequestError("email already exists")
	}

	return errors_utils.NewInternalServerError(fmt.Sprintf("error processing request"))
}
