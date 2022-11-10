package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"go.uber.org/zap"
	"log"
	"monorepo/src/auth_service/configs"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

// Initialize database connection then connect with postgres
func Init(config configs.Configuration) (*sqlx.DB, error) {

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDatabase)
	m, err := migrate.New("file://pkg/db/migrations", dbURL)
	if err != nil {
		log.Fatal("error in creating migrations: ", zap.Error(err))
	}
	fmt.Printf("")
	if err := m.Up(); err != nil {
		log.Println("error updating migrations: ", zap.Error(err))
	}

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
		// panic(err)
	}

	return db, nil
}
