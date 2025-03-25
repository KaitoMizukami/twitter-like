package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/KaitoMizukami/twitter-like/models"
)

type SqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepo(db *sql.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}

func (sr *SqliteRepo) GetAllPosts() ([]*models.Post, error) {
	rows, err := sr.db.Query("SELECT id, content, createdAt FROM posts;")
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
