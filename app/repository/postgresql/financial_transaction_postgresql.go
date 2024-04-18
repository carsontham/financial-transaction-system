package repository

import (
	"database/sql"
	"errors"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/repository"
	"financial-transaction-system/app/repository/dbmodel"
	"log"

	_ "github.com/lib/pq"
)

var _ repository.FinancialTransactionRepository = (*FinancialTransactionPostgresRepository)(nil)

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
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return dbmodel.AccountDBModelToDomainModel(account), nil
}

func (ftr FinancialTransactionPostgresRepository) CreateNewAccount(account *dbmodel.Account) error {
	db := ftr.getDB()
	if db == nil {
		log.Println("database error")
	}

	result, err := db.Exec("INSERT INTO account (id, balance) VALUES ($1, $2)",
		account.AccountID, account.Balance)
	if err != nil {
		return err
	}
	id, _ := result.RowsAffected()
	log.Printf("%d new row created: ", id)
	return nil
}

func (ftr FinancialTransactionPostgresRepository) CreateNewTransaction() error {
	postgresqlDB := ftr.getDB
	if postgresqlDB == nil {
		log.Println("database error")
	}
	return nil
}
