package storage

import (
	"fmt"
	"monorepo/src/customer_service/storage/postgres"
	"monorepo/src/customer_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Customer() repo.ICustomerStorage
}

type storagePg struct {
	customerRepo repo.ICustomerStorage
}

func New(db *sqlx.DB) IStorage {
	fmt.Println("storage ")
	storage := &storagePg{
		customerRepo: postgres.New(db),
	}
	return storage
}

func (s storagePg) Customer() repo.ICustomerStorage {
	return s.customerRepo
}
