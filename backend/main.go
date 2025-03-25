package main

import (
	"log"
	"net/http"

	"github.com/KaitoMizukami/twitter-like/db"
	"github.com/KaitoMizukami/twitter-like/handlers"
	"github.com/KaitoMizukami/twitter-like/migrations"
	"github.com/KaitoMizukami/twitter-like/repository"
	"github.com/KaitoMizukami/twitter-like/services"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	db, err := db.OpenSQLite("./twitterlike.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	err = migrations.CreatePostsTable(db)
	if err != nil {
		log.Fatal(err)
	}

	sr := repository.NewSqliteRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	ps := services.NewPostService(sr)
	handlers := handlers.NewPostHandler(ps)

	r.Get("/posts", handlers.GetAllPosts)

	http.ListenAndServe(":8080", r)
}
