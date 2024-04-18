package repository

import (
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/repository/dbmodel"
)

// FinancialTransactionRepository is the dependency to be injected into components that require access to the DB
type FinancialTransactionRepository interface {
	GetAccountByID(id int64) (*domain.Account, error)
	CreateNewAccount(account *dbmodel.Account) error
	CreateNewTransaction() error
}
