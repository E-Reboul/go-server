package server

import (
	"github.com/E-Reboul/go-server/routes/healthcheck"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	// Define group of healthcheck routes
	health := router.Group("/healthcheck")
	// Used groupd in healthcheck router to set all routes
	healthcheck.Router(health)
}
