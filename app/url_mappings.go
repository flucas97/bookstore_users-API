package app

import (
	c "github.com/flucas97/bookstore/users-api/controllers"
)

func MapUrl() {
	router.GET("/ping", c.Ping)

	router.POST("/users", c.CreateUser)
	router.GET("/user/:user_id", c.FindUser)
	router.GET("/users/search", c.SearchUser)
}
