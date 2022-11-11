package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type customerRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *customerRepo {
	return &customerRepo{db: db}
}
