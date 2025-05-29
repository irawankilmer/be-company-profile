package bootstrap

import (
	"company-profile/internal/config"
	"company-profile/internal/repository"
	"company-profile/internal/service"
)

type AppContainer struct {
	AuthService service.AuthService
	PostService service.PostService
}

func InitApp() *AppContainer {
	db := config.DB

	authRepo := repository.NewAuthRepository(db)
	postRepo := repository.NewPostRepository(db)

	return &AppContainer{
		AuthService: service.NewAuthService(authRepo),
		PostService: service.NewPostService(postRepo),
	}
}
