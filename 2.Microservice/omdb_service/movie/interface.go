package movie

import (
	"net/http"

	"omdb_service/infrastructure/omdb"
)

type Service interface {
	ProceedGetMovie(movieId, title string) (omdb.MovieInformation, error)
	ProceedGetMovies(searchWord, pagination string) (omdb.Result, error)
}

type Repository interface {
	InsertLog(log LogMovieSearch) error
}

type Middleware interface {
	LogSearch(next http.Handler) http.Handler
}
