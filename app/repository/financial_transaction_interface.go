package repository

import (
	"financial-transaction-system/app/domain"
)

//go:generate mockgen -source=financial_transaction_interface.go -package repositorytest -destination ../../tests/repositorytest/financial_transaction_repo_mock.go

// FinancialTransactionRepository is the dependency to be injected into components that require access to the DB
type FinancialTransactionRepository interface {
	GetAccountByID(id int64) (*domain.Account, error)
	CreateNewAccount(account *domain.Account) error
	GetTransactionByIdempotencyKey(string) (*domain.Transaction, error)
	PerformTransaction(*domain.Transaction) error
	GetAllTransactions() ([]*domain.Transaction, error)
}
