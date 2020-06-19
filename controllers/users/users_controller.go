package users_controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/flucas97/bookstore/users-api/model/users"
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
		err := utils.NewBadRequestError("ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, FindErr := services.FindUser(userID)
	if FindErr != nil {
		c.JSON(FindErr.Status, FindErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	// receive request, convert ID in int
	var err *utils.RestErr
	var userUpdates users.User

	userID, UserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if UserErr != nil {
		err = utils.NewBadRequestError("ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	// current user in database
	current, FindErr := services.FindUser(userID)
	if FindErr != nil {
		c.JSON(FindErr.Status, FindErr)
		return
	}

	// check if is put or patch
	if c.Request.Method == http.MethodPatch {
		userUpdates = *current
	}

	// fix created at blank after update
	userUpdates.CreatedAt = current.CreatedAt

	if err := c.ShouldBindJSON(&userUpdates); err != nil {
		restErr := utils.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, err.Error())
		return
	}

	userUpdates.ID = current.ID

	// persist changes
	result, err := services.UpdateUser(userUpdates)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	// status ok, user updated
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userID, UserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if UserErr != nil {
		err := utils.NewBadRequestError("ID should be a number")
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
		deleteErr := utils.NewInternalServerError(fmt.Sprintf("error while deleting user %v", user.ID))
		c.JSON(deleteErr.Status, deleteErr)
	}

	c.JSON(http.StatusOK, "user successfuly deleted")
}
