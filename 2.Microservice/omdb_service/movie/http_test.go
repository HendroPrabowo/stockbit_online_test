package movie

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"omdb_service/infrastructure/omdb"
)

type HttpTestSuite struct {
	suite.Suite
	getMovieUrl  string
	getMoviesUrl string
	sut          *Http
	service      *MockService
}

func (h *HttpTestSuite) SetupTest() {
	ctrl := gomock.NewController(h.T())
	h.getMovieUrl = "/movie?id=tt1672723&title=Batman"
	h.getMoviesUrl = "/movies?searchword=Batman&pagination=1"
	h.service = NewMockService(ctrl)
	h.sut = NewHttp(h.service)
}

func (h HttpTestSuite) TestNewHttp() {
	actual := NewHttp(h.service)
	h.Assert().Equal(h.sut, actual)
}

func (h HttpTestSuite) TestGetMovieHandlerSuccess() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, h.getMovieUrl, nil)
	h.service.EXPECT().ProceedGetMovie(gomock.Any(), gomock.Any()).Return(omdb.MovieInformation{}, nil)
	h.sut.GetMovieHandler(w, r)
	h.Assert().Equal(http.StatusOK, w.Code)
}

func (h HttpTestSuite) TestGetMovieHandlerGetErrorMovieNotFound() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, h.getMovieUrl, nil)
	h.service.EXPECT().ProceedGetMovie(gomock.Any(), gomock.Any()).Return(omdb.MovieInformation{}, errors.New(omdb.ErrMessageMovieNotFound))
	h.sut.GetMovieHandler(w, r)
	h.Assert().Equal(http.StatusBadRequest, w.Code)
}

func (h HttpTestSuite) TestGetMovieHandlerGetErrorInternalServer() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, h.getMovieUrl, nil)
	h.service.EXPECT().ProceedGetMovie(gomock.Any(), gomock.Any()).Return(omdb.MovieInformation{}, errors.New("internal server error"))
	h.sut.GetMovieHandler(w, r)
	h.Assert().Equal(http.StatusInternalServerError, w.Code)
}

func (h HttpTestSuite) TestGetMoviesHandlerSuccess() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, h.getMoviesUrl, nil)
	h.service.EXPECT().ProceedGetMovies(gomock.Any(), gomock.Any()).Return(omdb.Result{}, nil)
	h.sut.GetMoviesHandler(w, r)
	h.Assert().Equal(http.StatusOK, w.Code)
}

func (h HttpTestSuite) TestGetMoviesHandlerGetErrorMovieNotFound() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, h.getMoviesUrl, nil)
	h.service.EXPECT().ProceedGetMovies(gomock.Any(), gomock.Any()).Return(omdb.Result{}, errors.New(omdb.ErrMessageMovieNotFound))
	h.sut.GetMoviesHandler(w, r)
	h.Assert().Equal(http.StatusBadRequest, w.Code)
}

func (h HttpTestSuite) TestGetMoviesHandlerGetErrorInternalServer() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, h.getMoviesUrl, nil)
	h.service.EXPECT().ProceedGetMovies(gomock.Any(), gomock.Any()).Return(omdb.Result{}, errors.New("internal server error"))
	h.sut.GetMoviesHandler(w, r)
	h.Assert().Equal(http.StatusInternalServerError, w.Code)
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
