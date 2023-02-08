package db

import (
	config2 "advocate-back/pkg/config"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type postgres struct {
	p *sqlx.DB
}

func (p *postgres) DB() *sqlx.DB {
	return p.p
}

type DB interface {
	DB() *sqlx.DB
}

func NewPostgres() (*postgres, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config2.AppConfig.Postgres.Username,
		config2.AppConfig.Postgres.Password,
		config2.AppConfig.Postgres.Host,
		config2.AppConfig.Postgres.Port,
		config2.AppConfig.Postgres.DB,
	)
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &postgres{p: db}, nil
}
