package seeder

import (
	"company-profile/internal/config"
	"company-profile/internal/domain"
	"company-profile/pkg/utils"
)

func SeedAdminUser() {
	var adminRole domain.Role
	if err := config.DB.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		adminRole = domain.Role{Name: "admin"}
		config.DB.Create(&adminRole)
	}

	var adminUser domain.User
	if err := config.DB.Where("email = ?", "admin@example.com").First(&adminUser).Error; err == nil {
		return
	}

	hashed, _ := utils.HashPassword("admin123")

	adminUser = domain.User{
		Name:     "Administrator",
		Username: "admin",
		Email:    "admin@example.com",
		Password: hashed,
		Roles:    []domain.Role{adminRole},
	}

	config.DB.Create(&adminUser)
}
