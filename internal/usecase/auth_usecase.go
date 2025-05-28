package usecase

import (
	"company-profile/internal/config"
	"company-profile/internal/domain"
	"company-profile/pkg/utils"
	"errors"
)

func RegisterUser(user *domain.User) error {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	var role domain.Role
	config.DB.First(&role, "name = ?", "user") // default role, engke dirubah
	user.Roles = []domain.Role{role}

	return config.DB.Create(user).Error
}

func LoginUser(login, password string) (*domain.User, error) {
	var user domain.User
	config.DB.Preload("Roles").
		Where("email = ? OR username = ?", login, login).
		First(&user)

	if user.ID == 0 {
		return nil, errors.New("Invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}
