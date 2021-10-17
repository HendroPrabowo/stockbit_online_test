package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"omdb_service/infrastructure/config"
	"omdb_service/utils/httpUtils"
)

type ServiceImpl struct {
	url    string
	apikey string
}

func NewService(config config.AppConfig) Service {
	return &ServiceImpl{config.Omdb.Url, config.Omdb.ApiKey}
}

func (svc ServiceImpl) GetMovies(searchWord, pagination string) (result Result, err error) {
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%s", svc.url, svc.apikey, searchWord, pagination)
	resp, err := httpUtils.ForwardRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}
	if result.Response == FalseResponse {
		return result, ErrMovieNotFound
	}
	return
}

func (svc ServiceImpl) GetMovie(filter Filter) (movieInfo MovieInformation, err error) {
	url := fmt.Sprintf("%s/?apikey=%s&i=%s", svc.url, svc.apikey, filter.ID)
	if filter.Title != "" {
		url = fmt.Sprintf("%s/?apikey=%s&t=%s", svc.url, svc.apikey, filter.Title)
	}
	resp, err := httpUtils.ForwardRequest(http.MethodGet, url, nil)
	if err != nil {
		return movieInfo, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&movieInfo); err != nil {
		return movieInfo, err
	}
	if movieInfo.Response == FalseResponse {
		return movieInfo, ErrMovieNotFound
	}
	return
}
