package usecase

import (
	"company-profile/internal/dto/request"
	"company-profile/internal/repository"
	"company-profile/pkg/utils"
	"errors"
)

type AuthUsecase interface {
	Login(req request.LoginRequest) (string, error)
	Logout(token string) error
}

type authUsecase struct {
	repo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &authUsecase{repo}
}

func (s *authUsecase) Login(req request.LoginRequest) (string, error) {
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

func (s *authUsecase) Logout(token string) error {
	return nil
}
