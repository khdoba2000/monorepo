package storage

import (
	"monorepo/src/auth_service/storage/postgres"
	"monorepo/src/auth_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Authenitication() repo.IAuthStorage
}

type storagePg struct {
	authRepo repo.IAuthStorage
}

func New(db *sqlx.DB) *storagePg {
	return &storagePg{
		authRepo: postgres.New(db),
	}
}

func (s storagePg) Authenitication() repo.IAuthStorage {
	return s.authRepo
}
