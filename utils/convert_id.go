package utils

import (
	"fmt"
	"strconv"
)

func ConvertID(ID string) (int64, *RestErr) {
	userID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return 0, NewBadRequestError(fmt.Sprintf("invalid ID"))
	}

	return userID, nil
}
