package app

import (
	"github.com/gin-gonic/gin"
	"github.com/voicurobert/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	logger.Info("about to start the application")
	mapUrls()
	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}
}
