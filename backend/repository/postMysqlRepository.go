package repository

import (
	"database/sql"

	"github.com/KaitoMizukami/twitter-like/models"
)

type postMysqlRepository struct {
	db *sql.DB
}

func NewPostMysqlRepository(db *sql.DB) *postMysqlRepository {
	return &postMysqlRepository{
		db: db,
	}
}

func (mr *postMysqlRepository) GetAllPosts() ([]*models.Post, error) {
	rows, err := mr.db.Query("SELECT id, content, createdAt FROM posts;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
