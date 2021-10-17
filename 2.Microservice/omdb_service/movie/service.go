package movie

import "omdb_service/infrastructure/omdb"

type ServiceImpl struct {
	omdb omdb.Service
}

func NewService(omdb omdb.Service) Service {
	return &ServiceImpl{omdb}
}

func (svc ServiceImpl) ProceedGetMovie(movieId, title string) (omdb.MovieInformation, error) {
	filter := omdb.Filter{movieId, title}
	return svc.omdb.GetMovie(filter)
}

func (svc ServiceImpl) ProceedGetMovies(searchWord, pagination string) (omdb.Result, error) {
	return svc.omdb.GetMovies(searchWord, pagination)
}