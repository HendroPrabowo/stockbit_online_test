package movie

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"omdb_service/utils/logging"
)

type MiddlewareImpl struct {
	repository Repository
}

func NewMiddleware(repository Repository) Middleware {
	return &MiddlewareImpl{repository}
}

func (m MiddlewareImpl) LogSearch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newWriter := &logging.Recorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(newWriter, r)
		go m.saveLog(newWriter, r)
	})
}

func (m MiddlewareImpl) saveLog(w *logging.Recorder, r *http.Request) {
	log := LogMovieSearch{
		URL:        r.URL.Path,
		QueryParam: r.URL.RawQuery,
		StatusCode: w.StatusCode,
		CreatedAt:  time.Now(),
	}
	if w.Response != nil {
		log.Response = string(w.Response)
	}
	if err := m.repository.InsertLog(log); err != nil {
		logrus.Error(err)
	}
}
