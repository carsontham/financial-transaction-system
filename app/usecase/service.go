package usecase

import "financial-transaction-system/app/repository"

type Service struct {
	financialTransactionRepo repository.FinancialTransactionRepository
}

func NewService(repository repository.FinancialTransactionRepository) *Service {
	return &Service{
		financialTransactionRepo: repository,
	}
}
