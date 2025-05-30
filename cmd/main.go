package main

import (
	"company-profile/internal/bootstrap"
	"company-profile/internal/router"
	"github.com/gin-gonic/gin"
	"os"
)

// Swagger documentation
// @title Be Blog - REST API Docs
// @description Simply blog system
// @version 1.0
// @BasePath /
// @schemes http
// @schemes https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	app := bootstrap.InitApp()

	r := gin.Default()

	// Setup router
	router.SetupRoutes(r, app)

	r.Run(":" + port)
}
