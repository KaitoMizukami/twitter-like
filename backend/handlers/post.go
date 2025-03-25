package handlers

import (
	"encoding/json"
	"net/http"

	service "github.com/KaitoMizukami/twitter-like/services"
)

type postHandler struct {
	ps service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	return &postHandler{
		ps: service,
	}
}

func (ph *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.ps.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(posts)
}
