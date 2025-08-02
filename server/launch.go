package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Launch() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":66502"
	}
	// Initialize gin server with default configuration
	router := gin.Default()
	// Set up principal router to used in server
	Router(router)
	// If an discret error was occured on launch stop execution with error
	if err := router.Run(); err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}

// router.GET("/ping", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// })
