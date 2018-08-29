package main

import (
	"fmt"
	"wyatt/api/middleware"
	"wyatt/config"

	"wyatt/api/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// Disable Console Color
	gin.DisableConsoleColor()

	//server := gin.Default()
	// Creates a gin router with default middleware:
	server := gin.New()

	// logger and recovery (crash-free) middleware
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	server.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	server.Use(gin.Recovery())

	server.Use(middleware.Auth)

	server.Use(middleware.MustLogin)

	router.Router(server)

	// By default it serves on :8080 unless a
	//router.Run()
	port := config.GetConfig().Server.Port
	server.Run(fmt.Sprintf(":%d", port))
}
