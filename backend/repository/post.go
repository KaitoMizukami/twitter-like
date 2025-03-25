package repository

import "github.com/KaitoMizukami/twitter-like/models"

type PostRepository interface {
	GetAllPosts() ([]*models.Post, error)
}
