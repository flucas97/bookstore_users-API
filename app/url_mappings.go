package app

import (
	ping_controller "github.com/flucas97/bookstore/users-api/controllers/ping"
	users_controller "github.com/flucas97/bookstore/users-api/controllers/users"
)

func MapURL() {
	// ping route
	router.GET("/ping", ping_controller.Ping)

	// users routes
	router.GET("/user/:user_id", users_controller.Find)
	router.GET("/internal/users/search", users_controller.Search)

	router.POST("/users", users_controller.Create)
	router.POST("/users/login", users_controller.Login)

	router.PATCH("/user/:user_id", users_controller.Update)
	router.PUT("/user/:user_id", users_controller.Update)

	router.DELETE("/user/:user_id", users_controller.Delete)
}
