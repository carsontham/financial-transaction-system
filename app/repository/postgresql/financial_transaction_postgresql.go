package repository

import (
	"database/sql"
	"errors"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/repository"
	"financial-transaction-system/app/repository/dbmodel"
	"fmt"
	"github.com/lib/pq"
	"log"

	_ "github.com/lib/pq"
)

var _ repository.FinancialTransactionRepository = (*FinancialTransactionPostgresRepository)(nil)

type GetDB func() *sql.DB

type FinancialTransactionPostgresRepository struct {
	database *sql.DB
}

func NewFinancialTransactionRepository(db *sql.DB) *FinancialTransactionPostgresRepository {
	return &FinancialTransactionPostgresRepository{
		database: db,
	}
}

func (ftr FinancialTransactionPostgresRepository) GetAccountByID(id int64) (*domain.Account, error) {
	db := ftr.database
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

func (ftr FinancialTransactionPostgresRepository) CreateNewAccount(account *domain.Account) error {
	db := ftr.database
	if db == nil {
		log.Println("database error")
	}
	accDBModel := dbmodel.AccountDomainModelToDBModel(account)
	result, err := db.Exec("INSERT INTO account (id, balance) VALUES ($1, $2)",
		accDBModel.AccountID, accDBModel.Balance)

	if err != nil {
		return err
	}
	num, _ := result.RowsAffected()
	log.Printf("%d new row created: ", num)
	return nil
}

func (ftr FinancialTransactionPostgresRepository) GetTransactionByIdempotencyKey(key string) (*domain.Transaction, error) {
	db := ftr.database
	if db == nil {
		log.Println("database error")
	}

	query := "SELECT * FROM transaction WHERE idempotency_key = $1"
	row := db.QueryRow(query, key)

	var transaction dbmodel.Transaction
	err := row.Scan(
		&transaction.TransactionID,
		&transaction.SourceAccountID,
		&transaction.DestinationAccountID,
		&transaction.Amount, &transaction.IdempotencyKey,
	)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return dbmodel.TransactionDBModelToDomainModel(transaction), nil
}

func (ftr FinancialTransactionPostgresRepository) PerformTransaction(transaction *domain.Transaction) error {
	transactionDBModel := dbmodel.TransactionDomainModelToDBModel(transaction)
	db := ftr.database
	if db == nil {
		log.Println("database error")
	}

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			txn.Rollback()
			return
		}
		err = txn.Commit()
	}()

	if err := DeductAmount(db, transactionDBModel); err != nil {
		return err
	}
	if err := AddAmount(db, transactionDBModel); err != nil {
		return err
	}
	if err := CreateTransactionLog(db, transactionDBModel); err != nil {
		return err
	}
	return nil
}

func DeductAmount(db *sql.DB, transaction *dbmodel.Transaction) error {

	query := `UPDATE account SET balance = balance - $1 WHERE id = $2`
	result, err := db.Exec(query, transaction.Amount, transaction.SourceAccountID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			// "23514" is the error code for check constraint violation
			if pgErr.Code == "23514" {
				return domain.ErrInsufficientBalance
			}
		}
		return err
	}
	res, _ := result.RowsAffected()
	log.Printf("deducted amount for %d row", res)

	return nil
}

func AddAmount(db *sql.DB, transaction *dbmodel.Transaction) error {
	query := `UPDATE account SET balance = balance + $1 WHERE id = $2`
	result, err := db.Exec(query, transaction.Amount, transaction.DestinationAccountID)
	if err != nil {
		return err
	}
	res, _ := result.RowsAffected()
	log.Printf("added amount for %d row", res)
	return nil
}

func CreateTransactionLog(db *sql.DB, transaction *dbmodel.Transaction) error {
	query := `INSERT INTO transaction (source_account_id, destination_account_id, amount, idempotency_key)
				VALUES ($1, $2, $3, $4)`

	result, err := db.Exec(query, transaction.SourceAccountID, transaction.DestinationAccountID,
		transaction.Amount, transaction.IdempotencyKey)
	if err != nil {
		return err
	}
	res, _ := result.RowsAffected()
	log.Printf("inserted %d row into transaction table", res)
	return nil
}

func (ftr FinancialTransactionPostgresRepository) GetAllTransactions() ([]*domain.Transaction, error) {
	db := ftr.database
	if db == nil {
		log.Println("database error")
	}
	var transactions []*domain.Transaction
	rows, err := db.Query("SELECT * FROM transaction")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var transaction dbmodel.Transaction
		err := rows.Scan(&transaction.TransactionID, &transaction.SourceAccountID,
			&transaction.DestinationAccountID, &transaction.Amount, &transaction.IdempotencyKey)
		if err != nil {
			fmt.Println("error in scan", err)
			return nil, err
		}
		txn := dbmodel.TransactionDBModelToDomainModel(transaction)
		transactions = append(transactions, txn)
	}
	return transactions, nil
}
