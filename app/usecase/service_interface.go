package usecase

import (
	"financial-transaction-system/app/domain"
)

//go:generate mockgen -source=service_interface.go -package servicetest -destination ../../tests/servicetest/service_mock.go

type FinancialTransactionService interface {
	GetAccountById(int64) (*domain.Account, error)
	CreateNewAccount(*domain.Account) error
	PerformTransaction(txn *domain.Transaction) error
	GetAllTransactions() ([]*domain.Transaction, error)
	GetTransactionByIdempotencyKey(key string) (*domain.Transaction, error)
}
