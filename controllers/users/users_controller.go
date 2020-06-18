package users_controller

import (
	"log"
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
	// Aqui vai receber uma requisição e montar os dados
	// o ID do path
	userID, UserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if UserErr != nil {
		err := utils.NewBadRequestError("ID should be a number")
		c.JSON(err.Status, err)
		return
	}

	// com esse ID eu pesquiso o user no banco
	user, FindErr := services.FindUser(userID)
	if FindErr != nil {
		c.JSON(FindErr.Status, FindErr)
		return
	}
	log.Print(">>>>>>>>>>>>>>>>")
	// nesse user eu tenho que inserir o body do payload no user
	if err := c.ShouldBindJSON(user); err != nil {
		restErr := utils.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, err.Error())
		return
	}

	// aqui o usuário é de fato atualizado e salvo no banco
	if err := services.UpdateUser(user); err != nil {
		restErr := utils.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// retornamos status ok e o user foi atualizado
	c.JSON(http.StatusOK, user)
}
