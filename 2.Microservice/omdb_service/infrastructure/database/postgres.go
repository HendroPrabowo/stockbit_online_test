package database

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"

	"omdb_service/infrastructure/config"
)

func NewPostgres(config config.AppConfig) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     config.Postgres.Addr,
		User:     config.Postgres.User,
		Password: config.Postgres.Password,
		Database: config.Postgres.Database,
	})
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("cannot connect to postgres : %v", err.Error())
	}
	return db
}
