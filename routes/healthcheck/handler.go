package healthcheck

/*
 * Healthcheck handler for the API server.
 * This file contains the HealthcheckHandler function, which responds to healthcheck requests.
 */

import (
	"github.com/gin-gonic/gin"
)

// HealthcheckHandler responds to healthcheck request
func HealthcheckHandler(c *gin.Context) {
	// Check if has discret error
	if c.Request.Context().Err() != nil {
		c.String(500, "Internal Server Error")
		return
	}
	// Check HTTP method used to call this handler else return status error 405
	if c.Request.Method != "GET" {
		c.String(405, "Method Not Allowed")
		return
	}
	// If method is valid return status 200 else return error
	c.String(200, "Healthcheck ok")
}
