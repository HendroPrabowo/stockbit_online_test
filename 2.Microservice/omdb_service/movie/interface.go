package movie

import "omdb_service/infrastructure/omdb"

type Service interface {
	ProceedGetMovie(movieId, title string) (omdb.MovieInformation, error)
	ProceedGetMovies(searchWord, pagination string) (omdb.Result, error)
}
