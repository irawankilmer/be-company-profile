package bootstrap

import (
	"company-profile/internal/config"
	"company-profile/internal/repository"
	"company-profile/internal/seeder"
	"company-profile/internal/service"
	"os"
)

type AppContainer struct {
	AuthService     service.AuthService
	PostService     service.PostService
	CategoryService service.CategoryService
}

func InitApp() *AppContainer {
	config.LoadEnv()
	config.ConnectDB()
	db := config.DB

	if os.Getenv("APP_ENV") == "local" {
		seeder.SeedRoles()
		seeder.SeedAdminUser()
	}

	authRepo := repository.NewAuthRepository(db)
	postRepo := repository.NewPostRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	return &AppContainer{
		AuthService:     service.NewAuthService(authRepo),
		PostService:     service.NewPostService(postRepo),
		CategoryService: service.NewCategoryService(categoryRepo),
	}
}
