package bootstrap

import (
	"company-profile/internal/config"
	"company-profile/internal/repository"
	"company-profile/internal/service"
)

type AppContainer struct {
	PostService service.PostService
}

func InitApp() *AppContainer {
	db := config.DB

	postRepo := repository.NewPostRepository(db)

	return &AppContainer{
		PostService: service.NewPostService(postRepo),
	}
}
