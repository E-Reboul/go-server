package server

import (
	"os"

	logs "github.com/E-Reboul/go-server/middlewares/logs"
	"github.com/gin-gonic/gin"
)

func Launch() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":66502"
	}

	logs.LoadLogger()
	logs.WriteInfo(logs.ServerCategory, "Logs middleware initialization")

	logs.WriteInfo(logs.ServerCategory, "Initialization of Gin server...")
	// Initialize gin server with default configuration
	router := gin.Default()
	logs.WriteInfo(logs.ServerCategory, "Initialization of Gin server completed.")
	logs.WriteInfo(logs.ServerCategory, "Initialization of Router...")
	// Set up main router to used in server
	Router(router)
	logs.WriteInfo(logs.ServerCategory, "Initialization of Router completed.")

	// If an discrete error has occured on launch, stop execution with error
	if err := router.Run(); err != nil {
		logs.CloseLoggers()
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}

// router.GET("/ping", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// })
