package repository

import (
	"company-profile/internal/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmailOrUsername(value string) (*domain.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindByEmailOrUsername(value string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Roles").
		Where("email = ? OR username = ?", value, value).
		First(&user).Error

	return &user, err
}
