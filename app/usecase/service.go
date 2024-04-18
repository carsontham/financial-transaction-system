package usecase

import (
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/repository"
	"fmt"
)

type Service struct {
	financialTransactionRepo repository.FinancialTransactionRepository
}

func NewService(repository repository.FinancialTransactionRepository) *Service {
	return &Service{
		financialTransactionRepo: repository,
	}
}

func (s *Service) GetAccountById(id int64) (*domain.Account, error) {
	fmt.Println("in use case layer")
	account, err := s.financialTransactionRepo.GetAccountByID(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
