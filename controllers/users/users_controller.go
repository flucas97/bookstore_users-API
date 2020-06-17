package users_controller

import (
	"net/http"
	"strconv"

	"github.com/flucas97/bookstore/users-api/domain/users"
	"github.com/flucas97/bookstore/users-api/services"
	"github.com/flucas97/bookstore/users-api/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	// the same way of readall and unmarshall, but shortcut
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := utils.NewBadRequestError("Invalid JSON body")

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
	userID, UserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if UserErr != nil {
		err := utils.NewBadRequestError("Invalid ID")
		c.JSON(err.Status, err)
		return
	}

	user, FindErr := services.FindUser(userID)
	if FindErr != nil {
		err := utils.NewNotFoundError("User not found")
		c.JSON(err.Status, err)
		return
	}
	c.String(http.StatusNotImplemented, "Implement me")
}
