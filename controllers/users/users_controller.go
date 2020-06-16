package users_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flucas97/bookstore/users-api/domain/users"
	"github.com/flucas97/bookstore/users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(bytes, &user)
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := services.CreateUser(user)
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
