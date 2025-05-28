package main

import (
	"company-profile/internal/config"
	"company-profile/internal/delivery/handler"
	"company-profile/internal/seeder"
	"os"
)

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
