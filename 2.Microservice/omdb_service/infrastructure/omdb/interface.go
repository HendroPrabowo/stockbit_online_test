package omdb

type Service interface {
	GetMovies(searchWord, pagination string) (Result, error)
	GetMovie(filter Filter) (MovieInformation, error)
}
