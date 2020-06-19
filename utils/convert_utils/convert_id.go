package convert_utils

import (
	"fmt"
	"strconv"

	"github.com/flucas97/bookstore/users-api/utils/errors_utils"
)

func ConvertID(ID string) (int64, *errors_utils.RestErr) {
	userID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return 0, errors_utils.NewBadRequestError(fmt.Sprintf("invalid ID"))
	}

	return userID, nil
}
