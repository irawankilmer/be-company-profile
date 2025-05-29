package bootstrap

import (
	"company-profile/internal/config"
	"company-profile/internal/repository"
	"company-profile/internal/service"
)

type AppContainer struct {
	AuthService     service.AuthService
	PostService     service.PostService
	CategoryService service.CategoryService
}

func InitApp() *AppContainer {
	db := config.DB

	authRepo := repository.NewAuthRepository(db)
	postRepo := repository.NewPostRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	return &AppContainer{
		AuthService:     service.NewAuthService(authRepo),
		PostService:     service.NewPostService(postRepo),
		CategoryService: service.NewCategoryService(categoryRepo),
	}
}
