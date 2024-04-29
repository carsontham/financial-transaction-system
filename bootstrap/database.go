package bootstrap

import (
	"database/sql"
	"log"
)

func GetDB() *sql.DB {
	// these values can be placed in .env files
	const driverName = "postgres"
	const dataSourceName = "postgresql://root:secret@localhost:5432/financial-transaction-db?sslmode=disable"

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	return db
}
