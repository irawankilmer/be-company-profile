package main

import (
	"company-profile/internal/config"
	"company-profile/internal/delivery/handler"
	"company-profile/internal/seeder"
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
	config.LoadEnv()
	config.ConnectDB()

	if os.Getenv("APP_ENV") == "local" {
		seeder.SeedRoles()
	}

	r := handler.SetupRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
