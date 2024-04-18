package postgresql_repository

import (
	"database/sql"
	"log"
)

type GetDB func() *sql.DB

type FinancialTransactionRepository struct {
	getDB GetDB
}

func NewFinancialTransactionRepository() *FinancialTransactionRepository {
	return &FinancialTransactionRepository{
		getDB: GetDBConnection,
	}
}

func GetDBConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/memo-db?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	return db
}

func (ftr FinancialTransactionRepository) GetAccountByID() {
	postgresqlDB := ftr.getDB
	if postgresqlDB == nil {
		log.Println("database error")
	}
}

func (ftr FinancialTransactionRepository) CreateNewAccount() {
	postgresqlDB := ftr.getDB
	if postgresqlDB == nil {
		log.Println("database error")
	}
}

func (ftr FinancialTransactionRepository) CreateNewTransaction() {
	postgresqlDB := ftr.getDB
	if postgresqlDB == nil {
		log.Println("database error")
	}
}
