package repository

import (
	"database/sql"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/repository/dbmodel"
	"log"

	_ "github.com/lib/pq"
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

func (ftr FinancialTransactionPostgresRepository) GetAccountByID(id int64) (*domain.Account, error) {
	db := ftr.getDB()

	if db == nil {
		log.Println("database error")
	}

	query := "SELECT * FROM account WHERE id = $1"
	row := db.QueryRow(query, id)

	var account dbmodel.Account
	err := row.Scan(&account.AccountID, &account.Balance)
	if err != nil {
		return nil, err
	}
	return dbmodel.AccountDBModelToDomainModel(account), nil
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
