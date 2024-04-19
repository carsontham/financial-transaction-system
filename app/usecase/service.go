package usecase

import (
	"errors"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/repository"
	"financial-transaction-system/app/repository/dbmodel"
	"log"
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
	log.Println("in use case layer")
	account, err := s.financialTransactionRepo.GetAccountByID(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// IsFirstRequest checks whether account has already been created (serves as a form of idempotency)
func (s *Service) IsFirstRequest(id int64) bool {
	_, err := s.GetAccountById(id)
	if errors.Is(err, domain.ErrNotFound) {
		log.Println("new creation request - ID not found")
		return true
	}
	return false
}

func (s *Service) CreateNewAccount(account *domain.Account) error {
	log.Println("in use case layer")

	accDBModel := dbmodel.AccountDomainModelToDBModel(account)
	if err := s.financialTransactionRepo.CreateNewAccount(accDBModel); err != nil {
		log.Println("error creating new account: ", err)
		return err
	}
	return nil
}
