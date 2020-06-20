package users_controller

import (
	"fmt"
	"net/http"

	"github.com/flucas97/bookstore/users-api/model/users"
	"github.com/flucas97/bookstore/users-api/services"
	"github.com/flucas97/bookstore/users-api/utils/convert_utils"
	"github.com/flucas97/bookstore/users-api/utils/errors_utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors_utils.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {
	userID, UserErr := convert_utils.ConvertID(c.Param("user_id"))
	if UserErr != nil {
		err := errors_utils.NewBadRequestError("ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, FindErr := services.FindUser(userID)
	if FindErr != nil {
		c.JSON(FindErr.Status, FindErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	var err *errors_utils.RestErr
	var userUpdates users.User

	userID, UserErr := convert_utils.ConvertID(c.Param("user_id"))
	if UserErr != nil {
		err = errors_utils.NewBadRequestError("ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	current, FindErr := services.FindUser(userID)
	if FindErr != nil {
		c.JSON(FindErr.Status, FindErr)
		return
	}

	if c.Request.Method == http.MethodPatch {
		userUpdates = *current
	}

	userUpdates.CreatedAt = current.CreatedAt

	if err := c.ShouldBindJSON(&userUpdates); err != nil {
		restErr := errors_utils.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, err.Error())
		return
	}

	userUpdates.ID = current.ID

	result, err := services.UpdateUser(userUpdates)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userID, UserErr := convert_utils.ConvertID(c.Param("user_id"))
	if UserErr != nil {
		err := errors_utils.NewBadRequestError("ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, FindErr := services.FindUser(userID)
	if FindErr != nil {
		c.JSON(FindErr.Status, FindErr)
		return
	}

	err := services.DeleteUser(user)
	if err != nil {
		deleteErr := errors_utils.NewInternalServerError(fmt.Sprintf("error while deleting user %v", user.ID))
		c.JSON(deleteErr.Status, deleteErr)
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, FindErr := services.Search(status)
	if FindErr != nil {
		c.JSON(FindErr.Status, FindErr)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
