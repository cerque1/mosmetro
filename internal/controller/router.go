package controller

import (
	"net/http"

	"github.com/go-chi/chi"

	my_http "mosmetro/internal/controller/http"
	"mosmetro/internal/usecase"
)

func NewRouter(us usecase.Stations) http.Handler {
	router := chi.NewRouter()
	router.Mount("/api", my_http.NewRouter(us))
	return router
}
