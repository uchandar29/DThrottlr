package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() http.Handler {
	router := chi.NewRouter()

	/* Router  */
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	/* Routes & Handler */
	router.Get("/health", HealthHandler)

	return router
}

func Run(address string) error {
	router := New()
	return http.ListenAndServe(address, router)
}
