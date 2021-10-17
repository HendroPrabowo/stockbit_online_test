package movie

import (
	"context"
	"errors"

	"omdb_service/infrastructure/config"
	"omdb_service/infrastructure/grpc"
	"omdb_service/infrastructure/omdb"
)

type Grpc struct {
	grpc.UnimplementedMovieServer
	service Service
}

func NewGrpc(config config.AppConfig) *Grpc {
	omdb := omdb.NewService(config)
	service := NewService(omdb)
	return &Grpc{service: service}
}

func (g Grpc) GetMovies(ctx context.Context, req *grpc.MoviesRequest) (*grpc.MovieResult, error) {
	result, err := g.service.ProceedGetMovies(req.Searchword, req.Pagination)
	if err != nil {
		return nil, err
	}
	if result.Response != omdb.FalseResponse {
		return nil, errors.New(result.Error)
	}
	return g.mapMoviesResult(result), nil
}

func (g Grpc) GetMovie(ctx context.Context, req *grpc.MovieRequest) (*grpc.MovieInformation, error) {
	result, err := g.service.ProceedGetMovie(req.Id, req.Title)
	if err != nil {
		return nil, err
	}
	if result.Response != omdb.FalseResponse {
		return nil, errors.New(result.Error)
	}
	return g.mapMovieResult(result), nil
}

func (g Grpc) mapMoviesResult(result omdb.Result) *grpc.MovieResult {
	response := &grpc.MovieResult{
		TotalResult: result.TotalResult,
		Response:    result.Response,
		Error:       result.Error,
	}
	for _, val := range result.Search {
		summary := &grpc.MovieSummary{
			Title:  val.Title,
			Year:   val.Year,
			Imdbid: val.ImdbID,
			Type:   val.Type,
			Poster: val.Poster,
		}
		response.MovieSummary = append(response.MovieSummary, summary)
	}
	return response
}

func (g Grpc) mapMovieResult(result omdb.MovieInformation) *grpc.MovieInformation {
	response := &grpc.MovieInformation{
		Title:      result.Title,
		Year:       result.Year,
		Rated:      result.Rated,
		Released:   result.Released,
		Runtime:    result.Runtime,
		Genre:      result.Genre,
		Director:   result.Director,
		Writter:    result.Writer,
		Actors:     result.Actors,
		Plot:       result.Plot,
		Language:   result.Language,
		Country:    result.Country,
		Awards:     result.Awards,
		Poster:     result.Poster,
		Metascore:  result.Metascore,
		Imdbrating: result.ImdbRating,
		Imdbvotes:  result.ImdbVotes,
		Imdbid:     result.ImdbID,
		Type:       result.Type,
		Dvd:        result.DVD,
		Boxoffice:  result.BoxOffice,
		Production: result.Production,
		Website:    result.Website,
		Response:   result.Response,
		Error:      result.Error,
	}
	for _, val := range result.Ratings {
		rating := &grpc.Rating{
			Source: val.Source,
			Value:  val.Value,
		}
		response.Ratings = append(response.Ratings, rating)
	}
	return response
}
