package seeder

import (
	"company-profile/internal/config"
	"company-profile/internal/domain"
)

func SeedRoles() {
	roles := []domain.Role{
		{Name: "admin"},
		{Name: "user"},
		{Name: "moderator"},
	}

	for _, r := range roles {
		config.DB.FirstOrCreate(&r, domain.Role{Name: r.Name})
	}
}
