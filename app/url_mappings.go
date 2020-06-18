package app

import (
	ping_controller "github.com/flucas97/bookstore/users-api/controllers/ping"
	users_controller "github.com/flucas97/bookstore/users-api/controllers/users"
)

func MapURL() {
	router.GET("/ping", ping_controller.Ping)

	router.PATCH("/user/:user_id", users_controller.UpdateUser)
	router.POST("/users", users_controller.CreateUser)
	router.GET("/user/:user_id", users_controller.FindUser)
}
