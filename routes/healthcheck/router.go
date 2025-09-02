package healthcheck

/*
 * Healthcheck router configuration for the API server.
 * This file contains the Router function, which sets up the healthcheck route group and attaches its handler.
 */

import "github.com/gin-gonic/gin"

// Router sets up the healthcheck route group and attaches its handler.
func Router(routerGroup *gin.RouterGroup) {
	// Set up all methods allowed to healthcheck routerGroup
	routerGroup.GET("/", HealthcheckHandler)
}
