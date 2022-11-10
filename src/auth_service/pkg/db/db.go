package db

import (
	"fmt"
	"monorepo/src/auth_service/configs"

	"github.com/jmoiron/sqlx"
)

// Initialize database connection then connect with postgres
func Init(config configs.Configuration) (*sqlx.DB, error) {

	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDatabase,
	)

	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return nil, err
		// panic(err)
	}

	return db, nil
}
