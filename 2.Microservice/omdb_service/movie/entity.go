package movie

import "time"

type LogMovieSearch struct {
	tableName struct{} `pg:"log_movie_search"`
	ID         int
	URL        string
	QueryParam string
	StatusCode int
	Response   string
	CreatedAt  time.Time
}
