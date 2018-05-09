package postgres

import (
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Url string
}

func NewPostgres(cfg *Config) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", cfg.Url)
}
