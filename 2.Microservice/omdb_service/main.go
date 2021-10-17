package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"

	"omdb_service/infrastructure/config"
	"omdb_service/movie"
)

func init() {
	config.Init()
}

func main() {
	r := chi.NewRouter()
	setCors(r)
	r.Use(middleware.Logger)

	// inject movie routes
	movieRoutes := movie.NewRoute(config.Config)
	movieRoutes.RegisterRoutes(r)

	logrus.Infof("%s running on port %s\n", config.Config.ServiceName, config.Config.Server.Port)
	http.ListenAndServe(":"+config.Config.Server.Port, r)
}

func setCors(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}
