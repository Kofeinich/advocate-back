package db

import (
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
	db, err := sqlx.Connect("posts", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &postgres{p: db}, nil
}
