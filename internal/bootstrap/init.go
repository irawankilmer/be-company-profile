package bootstrap

import (
	"company-profile/internal/config"
	"company-profile/internal/migration"
	"company-profile/internal/repository"
	"company-profile/internal/seeder"
	"company-profile/internal/usecase"
	"os"
)

type AppContainer struct {
	AuthUsecase     usecase.AuthUsecase
	PostUsecase     usecase.PostUsecase
	CategoryUsecase usecase.CategoryUsecase
}

func InitApp() *AppContainer {
	config.LoadEnv()
	config.ConnectDB()
	db := config.DB

	if os.Getenv("APP_ENV") == "local" {
		migration.AutoMigrate()
		seeder.SeedRoles()
		seeder.SeedAdminUser()
	}

	authRepo := repository.NewAuthRepository(db)
	postRepo := repository.NewPostRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	return &AppContainer{
		AuthUsecase:     usecase.NewAuthUsecase(authRepo),
		PostUsecase:     usecase.NewPostUsecase(postRepo),
		CategoryUsecase: usecase.NewCategoryUsecase(categoryRepo),
	}
}
