package models

import "time"

type Post struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
