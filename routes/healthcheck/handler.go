package healthcheck

import (
	"github.com/gin-gonic/gin"
)

func HealthcheckGetHandler(c *gin.Context) {
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

// LOGIC WITHOUT GIN

// // Check if has discret error
// if err := request.Context().Err(); err != nil {
// 	response.WriteHeader(http.StatusInternalServerError)
// }

// // Check HTTP method used to call this handler else return status error 405
// if request.Method != http.MethodGet {
// 	response.WriteHeader(http.StatusMethodNotAllowed)
// 	_, _ = response.Write([]byte("Method Not Allowed"))
// 	return
// }

// // If method is valid return status 200 else return error
// response.WriteHeader(http.StatusOK)
// if _, err := response.Write([]byte("Healthcheck ok")); err != nil {
// 	fmt.Printf("Healthcheck failed: %v\n", err)
// }
