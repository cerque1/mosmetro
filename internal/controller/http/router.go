package http

import (
	"net/http"

	"github.com/go-chi/chi"

	v1 "mosmetro/internal/controller/http/v1"
	"mosmetro/internal/usecase"
)

func NewRouter(us usecase.Stations) http.Handler {
	router := chi.NewRouter()
	router.Mount("/v1", v1.NewRouter(us))
	return router
}
