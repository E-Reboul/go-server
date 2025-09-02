package server

/*
 * Main router configuration for the API server.
 * Router configures all route groups and attaches their handlers to the Gin engine
 * Extend this function to register new router modules.
 */
/*
 * Main router configuration for the API server.
 *
 * This file contains the Router function, which configures all route groups and attaches their handlers to the Gin engine.
 * Extend this file to register new router modules.
 */

import (
	"github.com/E-Reboul/go-server/routes/healthcheck"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	/*
	  ===========
	  Healthcheck
	  ===========
	*/

	health := router.Group("/healthcheck")
	healthcheck.Router(health)
}
