package movie

import (
	"github.com/go-chi/chi/v5"

	"omdb_service/infrastructure/config"
	"omdb_service/infrastructure/database"
	"omdb_service/infrastructure/omdb"
)

type Route struct {
	http       *Http
	middleware Middleware
}

func NewRoute(config config.AppConfig) *Route {
	omdbService := omdb.NewService(config)
	service := NewService(omdbService)
	http := NewHttp(service)
	postgres := database.NewPostgres(config)
	repository := NewRepository(postgres)
	middleware := NewMiddleware(repository)
	return &Route{http, middleware}
}

func (route Route) RegisterRoutes(r chi.Router) {
	r.Use(route.middleware.LogSearch)
	r.Get("/movie", route.http.GetMovieHandler)
	r.Get("/movies", route.http.GetMoviesHandler)
}
