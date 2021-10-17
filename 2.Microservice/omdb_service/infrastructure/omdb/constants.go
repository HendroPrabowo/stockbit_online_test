package omdb

import "errors"

const (
	FalseResponse = "False"
)

const (
	ErrMessageMovieNotFound   = "Movie not found!"
	ErrMessageIncorrectImdbID = "Incorrect IMDb ID."
)

var (
	ErrMovieNotFound = errors.New("Movie not found!")
)
