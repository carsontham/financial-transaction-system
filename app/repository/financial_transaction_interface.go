package repository

import "financial-transaction-system/app/domain"

// FinancialTransactionRepository is the dependency to be injected into components that require access to the DB
type FinancialTransactionRepository interface {
	GetAccountByID(id int64) (*domain.Account, error)
	CreateNewAccount() error
	CreateNewTransaction() error
}
