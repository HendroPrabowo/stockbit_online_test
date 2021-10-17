package movie

import (
	"github.com/go-chi/chi/v5"

	"omdb_service/infrastructure/config"
	"omdb_service/infrastructure/omdb"
)

type Route struct {
	Http *Http
}

func NewRoute(config config.AppConfig) *Route {
	omdbService := omdb.NewService(config)
	service := NewService(omdbService)
	http := NewHttp(service)
	return &Route{http}
}

func (route Route) RegisterRoutes(r chi.Router) {
	r.Get("/movie", route.Http.GetMovieHandler)
	r.Get("/movies", route.Http.GetMoviesHandler)
}
