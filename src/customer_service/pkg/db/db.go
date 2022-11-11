package db

import (
	"fmt"
	"monorepo/src/customer_service/configs"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

// Initialize database connection then connect with postgres
func Init(config *configs.Configuration) (*sqlx.DB, error) {

	fmt.Println("init db")
	// dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	config.PostgresHost,
	// 	config.PostgresPort,
	// 	config.PostgresUser,
	// 	config.PostgresPassword,
	// 	config.PostgresDatabase,
	// )

	// db, err := sqlx.Connect("postgres", dbUrl)
	// if err != nil {
	// 	panic(err)
	// }
	db := &sqlx.DB{}

	return db, nil
}
