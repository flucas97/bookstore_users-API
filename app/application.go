package app

import (
	"github.com/flucas97/bookstore/users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	log    = logger.Logger
)

func StartApplication() {
	MapURL()

	log.Info("starting application...")
	router.Run(":8080")
}
