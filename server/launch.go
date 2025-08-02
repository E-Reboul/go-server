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

	router := gin.Default()

	err := router.Run()

	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}

// router.GET("/ping", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// })
