package services

import (
	"github.com/KaitoMizukami/twitter-like/models"
	"github.com/KaitoMizukami/twitter-like/repository"
)

type PostService interface {
	GetAllPosts() ([]*models.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *postService {
	return &postService{
		repo: repo,
	}
}

func (ps *postService) GetAllPosts() ([]*models.Post, error) {
	return ps.repo.GetAllPosts()
}
