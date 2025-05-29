package repository

import (
	"company-profile/internal/domain"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *domain.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(post *domain.Post) error {
	return r.db.Create(post).Error
}
