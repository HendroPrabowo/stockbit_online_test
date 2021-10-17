package movie

import (
	"net/http"

	"github.com/go-chi/render"

	"omdb_service/infrastructure/omdb"
	"omdb_service/utils/response"
)

type Http struct {
	service Service
}

func NewHttp(service Service) *Http {
	return &Http{service}
}

func (h Http) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieId := r.URL.Query().Get(ID)
	title := r.URL.Query().Get(Title)
	movie, err := h.service.ProceedGetMovie(movieId, title)
	if err != nil {
		if err.Error() == omdb.ErrMessageIncorrectImdbID || err.Error() == omdb.ErrMessageMovieNotFound {
			response.HttpResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.HttpResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, movie)
}

func (h Http) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	pagination := r.URL.Query().Get(Pagination)
	searchWord := r.URL.Query().Get(SearchWord)
	movies, err := h.service.ProceedGetMovies(searchWord, pagination)
	if err != nil {
		if err.Error() == omdb.ErrMessageIncorrectImdbID || err.Error() == omdb.ErrMessageMovieNotFound {
			response.HttpResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.HttpResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, movies)
}
