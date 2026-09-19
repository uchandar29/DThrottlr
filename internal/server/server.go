package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/uchandar29/DThrottlr/internal/limiter"
	"github.com/uchandar29/DThrottlr/internal/middleware"
)

func New(limiter *limiter.Limiter) http.Handler {
	router := chi.NewRouter()

	/* Router  */
	router.Use(chimw.Logger)
	router.Use(chimw.Recoverer)
	router.Use(middleware.RateLimit(limiter))

	/* Routes & Handler */
	router.Get("/health", HealthHandler)
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return router
}

func Run(address string, limiter *limiter.Limiter) error {
	router := New(limiter)
	return http.ListenAndServe(address, router)
}
