package repository

import (
	"database/sql"
	"log"
)

type GetDB func() *sql.DB

type FinancialTransactionPostgresRepository struct {
	getDB GetDB
}

func NewFinancialTransactionRepository(getDBConnection GetDB) *FinancialTransactionPostgresRepository {
	return &FinancialTransactionPostgresRepository{
		getDB: getDBConnection,
	}
}

func (ftr FinancialTransactionPostgresRepository) GetAccountByID() error {
	postgresqlDB := ftr.getDB()
	if postgresqlDB == nil {
		log.Println("database error")
	}
	return nil
}

func (ftr FinancialTransactionPostgresRepository) CreateNewAccount() error {
	postgresqlDB := ftr.getDB
	if postgresqlDB == nil {
		log.Println("database error")
	}
	return nil
}

func (ftr FinancialTransactionPostgresRepository) CreateNewTransaction() error {
	postgresqlDB := ftr.getDB
	if postgresqlDB == nil {
		log.Println("database error")
	}
	return nil
}
