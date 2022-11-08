package storage

import (
	"github.com/jmoiron/sqlx"
	"monorepo/src/auth_service/storage/postgres"
	"monorepo/src/auth_service/storage/repo"
)

type IStorage interface {
	Authenitication() repo.IAuthStorage
}

type storagePg struct {
	db       *sqlx.DB
	authRepo repo.IAuthStorage
}

func New(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		authRepo: postgres.New(db),
	}
}

func (s storagePg) Authenitication() repo.IAuthStorage {
	return s.authRepo
}
