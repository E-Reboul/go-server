package healthcheck

import "github.com/gin-gonic/gin"

func Router(routerGroup *gin.RouterGroup) {
	// Set up all methods allowed to healthcheck routerGroup
	routerGroup.GET("/", HealthcheckHandler)
}
