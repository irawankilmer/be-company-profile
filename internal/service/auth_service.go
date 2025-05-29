package service

import (
	"company-profile/internal/dto/request"
	"company-profile/internal/repository"
	"company-profile/pkg/utils"
	"errors"
)

type AuthService interface {
	Login(req request.LoginRequest) (string, error)
	Logout(token string) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(req request.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmailOrUsername(req.Login)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("email/username atau password salah")
	}

	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	token, err := utils.GenerateJWT(user.ID, roleNames)
	return token, err
}

func (s *authService) Logout(token string) error {
	return nil
}
