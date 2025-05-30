package usecase

import (
	"company-profile/internal/domain"
	"company-profile/internal/dto/request"
	"company-profile/internal/repository"
)

type CategoryUsecase interface {
	Create(req request.CreateCategoryRequest) (*domain.Category, error)
	GetAll() ([]domain.Category, error)
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(r repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{r}
}

func (s *categoryUsecase) Create(req request.CreateCategoryRequest) (*domain.Category, error) {
	category := domain.Category{
		Name:             req.Name,
		Description:      req.Description,
		ParentCategoryID: req.ParentCategoryID,
	}
	err := s.repo.Create(&category)
	return &category, err
}

func (s *categoryUsecase) GetAll() ([]domain.Category, error) {
	return s.repo.GetAll()
}
