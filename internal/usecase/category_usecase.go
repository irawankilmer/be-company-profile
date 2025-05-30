package usecase

import (
	"company-profile/internal/domain"
	"company-profile/internal/dto/request"
	"company-profile/internal/repository"
)

type CategoryUsecase interface {
	Create(req request.CreateCategoryRequest) (*domain.Category, error)
	GetAll() ([]domain.Category, error)
	GetByID(id uint) (*domain.Category, error)
	Update(id uint, req request.UpdateCategoryRequest) (*domain.Category, error)
	Delete(id uint) error
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

func (s *categoryUsecase) GetByID(id uint) (*domain.Category, error) { return s.repo.GetByID(id) }

func (s *categoryUsecase) Update(id uint, req request.UpdateCategoryRequest) (*domain.Category, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}

	err = s.repo.Update(category)
	return category, err
}

func (s *categoryUsecase) Delete(id uint) error {
	category, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(category)
}
