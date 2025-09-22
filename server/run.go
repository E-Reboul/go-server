package server

/*
 * Initialization and run API server.
 * This file contain Run method including initialization of loggers and Gin Server
 */

import (
	"os"

	logs "github.com/E-Reboul/go-server/middlewares/logs"

	"github.com/gin-gonic/gin"
)

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":66502"
	}

	/*
	  ===========
	  Logs Server
	  ===========
	*/

	logs.InitializeLoggers()
	logs.WriteInfo(logs.ServerCategory, "Initialization of Logs middleware completed.")

	/*
	  ==========
	  Gin Server
	  ==========
	*/

	logs.WriteInfo(logs.ServerCategory, "Initialization of Gin server...")
	router := gin.Default()
	logs.WriteInfo(logs.ServerCategory, "Initialization of Gin server completed.")

	/*
	  =============
	  Global router
	  =============
	*/

	logs.WriteInfo(logs.ServerCategory, "Initialization of Router...")
	// Set up primary Router with they childs
	Router(router)
	logs.WriteInfo(logs.ServerCategory, "Initialization of Router completed.")

	// If an discrete error has occured on launch, stop execution with error
	if err := router.Run(); err != nil {
		logs.WriteError(logs.ServerCategory, "Server has failed to start: "+err.Error())
		logs.CloseLoggers()
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
