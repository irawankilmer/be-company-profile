package handler

import (
	"company-profile/docs"
	"company-profile/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// dokumentasi swagger
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", Register)
	r.POST("/login", Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/posts", GetPosts)
	auth.POST("/posts", middleware.RoleMiddleware("user", "admin"), CreatePost)

	return r
}
