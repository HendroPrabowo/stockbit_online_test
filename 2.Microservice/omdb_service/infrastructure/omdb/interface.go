package omdb

type Service interface {
	GetMovies(searchKey, pagination string) (Result, error)
	GetMovie(filter Filter) (MovieInformation, error)
}
