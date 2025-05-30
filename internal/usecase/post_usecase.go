package usecase

import (
	"company-profile/internal/domain"
	"company-profile/internal/dto/request"
	"company-profile/internal/repository"
)

type PostUsecase interface {
	Create(req request.CreatePostRequest, userID uint) (*domain.Post, error)
	GetAll() ([]domain.Post, error)
	GetByID(id uint) (*domain.Post, error)
	Update(id uint, req request.UpdatePostRequest) (*domain.Post, error)
	Delete(id uint) error
}

type postUsecase struct {
	repo repository.PostRepository
}

func NewPostUsecase(r repository.PostRepository) PostUsecase {
	return &postUsecase{r}
}

func (s *postUsecase) Create(req request.CreatePostRequest, userID uint) (*domain.Post, error) {
	post := domain.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}
	err := s.repo.Create(&post)
	return &post, err
}

func (s *postUsecase) GetAll() ([]domain.Post, error) {
	return s.repo.GetAll()
}

func (s *postUsecase) GetByID(id uint) (*domain.Post, error) {
	return s.repo.GetByID(id)
}

func (s *postUsecase) Update(id uint, req request.UpdatePostRequest) (*domain.Post, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}
	err = s.repo.Update(post)
	return post, err
}

func (s *postUsecase) Delete(id uint) error {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(post)
}
