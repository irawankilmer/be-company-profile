package handler

import (
	"company-profile/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", Register)
	r.POST("/login", Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/posts", GetPosts)
	auth.POST("/posts", middleware.RoleMiddleware("user", "admin"), CreatePost)

	return r
}
