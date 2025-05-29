package service

import (
	"company-profile/internal/domain"
	"company-profile/internal/dto/request"
	"company-profile/internal/repository"
)

type CategoryService interface {
	Create(req request.CreateCategoryRequest) (*domain.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{r}
}

func (s *categoryService) Create(req request.CreateCategoryRequest) (*domain.Category, error) {
	category := domain.Category{
		Name:             req.Name,
		Description:      req.Description,
		ParentCategoryID: req.ParentCategoryID,
	}
	err := s.repo.Create(&category)
	return &category, err
}
