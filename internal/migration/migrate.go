package migration

import (
	"company-profile/internal/config"
	"company-profile/internal/domain"
)

func AutoMigrate() {
	db := config.DB
	db.AutoMigrate(
		&domain.User{},
		&domain.Role{},
		&domain.Post{},
		&domain.Category{},
	)
}
