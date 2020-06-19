package app

import (
	ping_controller "github.com/flucas97/bookstore/users-api/controllers/ping"
	users_controller "github.com/flucas97/bookstore/users-api/controllers/users"
)

func MapURL() {
	// ping route
	router.GET("/ping", ping_controller.Ping)

	// users routes
	router.GET("/user/:user_id", users_controller.FindUser)
	router.GET("/internal/users/search", users_controller.Search)

	router.POST("/users", users_controller.CreateUser)

	router.PATCH("/user/:user_id", users_controller.UpdateUser)
	router.PUT("/user/:user_id", users_controller.UpdateUser)

	router.DELETE("/user/:user_id", users_controller.DeleteUser)
}
