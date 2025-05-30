package bootstrap

import (
	"company-profile/internal/config"
	"company-profile/internal/repository"
	"company-profile/internal/seeder"
	"company-profile/internal/usecase"
	"os"
)

type AppContainer struct {
	AuthService     usecase.AuthUsecase
	PostService     usecase.PostUsecase
	CategoryService usecase.CategoryUsecase
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
		AuthService:     usecase.NewAuthUsecase(authRepo),
		PostService:     usecase.NewPostUsecase(postRepo),
		CategoryService: usecase.NewCategoryUsecase(categoryRepo),
	}
}
