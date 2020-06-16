package app

import "github.com/flucas97/bookstore/users-api/controllers"

func MapUrl() {
	router.GET("/ping", controllers.Ping)
}
