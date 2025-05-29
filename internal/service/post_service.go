package service

import (
	"company-profile/internal/domain"
	"company-profile/internal/dto/request"
	"company-profile/internal/repository"
)

type PostService interface {
	Create(req request.CreatePostRequest, userID uint) (*domain.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{r}
}

func (s *postService) Create(req request.CreatePostRequest, userID uint) (*domain.Post, error) {
	post := domain.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}
	err := s.repo.Create(&post)
	return &post, err
}
