package movie

import "github.com/go-pg/pg/v10"

type RepositoryImpl struct {
	postgres *pg.DB
}

func NewRepository(postgres *pg.DB) Repository {
	return &RepositoryImpl{postgres}
}

func (repo RepositoryImpl) InsertLog(log LogMovieSearch) error {
	_, err := repo.postgres.Model(&log).Insert()
	return err
}
