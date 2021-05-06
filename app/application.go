package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
