package repository

import (
	"company-profile/internal/domain"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *domain.Post) error
	GetAll() ([]domain.Post, error)
	GetByID(id uint) (*domain.Post, error)
	Update(post *domain.Post) error
	Delete(post *domain.Post) error
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

func (r *postRepository) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *postRepository) GetByID(id uint) (*domain.Post, error) {
	var post domain.Post
	err := r.db.First(&post, id).Error
	return &post, err
}

func (r *postRepository) Update(post *domain.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(post *domain.Post) error {
	return r.db.Delete(post).Error
}
