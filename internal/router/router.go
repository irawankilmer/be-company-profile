package router

import (
	"company-profile/docs"
	"company-profile/internal/bootstrap"
	"company-profile/internal/handler"
	"company-profile/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, app *bootstrap.AppContainer) {
	r.Use(middleware.CORSMiddleware())

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authHandler := handler.NewAuthHandler(app.AuthService)
	postHandler := handler.NewPostHandler(app.PostService)
	categoryHandler := handler.NewCategoryHandler(app.CategoryService)

	api := r.Group("/api")
	{
		api.POST("/login", authHandler.Login)
		api.POST("/logout", authHandler.Logout)
	}

	admin := r.Group("/api").Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	admin.GET("/post", postHandler.GetAllPosts)
	admin.POST("/post", postHandler.CreatePost)
	admin.GET("/post/:id", postHandler.GetPostByID)
	admin.PUT("/post/:id", postHandler.UpdatePost)
	admin.DELETE("/post/:id", postHandler.DeletePost)

	admin.POST("/category", categoryHandler.CreateCategory)
}
