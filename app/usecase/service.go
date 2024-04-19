package usecase

import (
	"errors"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/repository"
	"log"
)

var _ FinancialTransactionService = (*Service)(nil)

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

// CheckIfAccountExist checks whether account exists (serves as a form of idempotency check)
func CheckIfAccountExist(service FinancialTransactionService, id int64) bool {
	_, err := service.GetAccountById(id)
	if errors.Is(err, domain.ErrNotFound) {
		log.Println("account  - ID not found")
		return false
	}
	return true
}

func (s *Service) CreateNewAccount(account *domain.Account) error {
	log.Println("in use case layer")

	if err := s.financialTransactionRepo.CreateNewAccount(account); err != nil {
		log.Println("error creating new account: ", err)
		return err
	}
	return nil
}

func (s *Service) GetTransactionByIdempotencyKey(key string) (*domain.Transaction, error) {
	transaction, err := s.financialTransactionRepo.GetTransactionByIdempotencyKey(key)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// PerformTransaction checks for valid source and destination accounts before making the transfer
func (s *Service) PerformTransaction(txn *domain.Transaction) error {
	isValidSourceAccount := CheckIfAccountExist(s, txn.SourceAccountID)
	isValidDestinationAccount := CheckIfAccountExist(s, txn.DestinationAccountID)
	if !isValidSourceAccount || !isValidDestinationAccount {
		return domain.ErrNotFound
	}

	err := s.financialTransactionRepo.PerformTransaction(txn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAllTransactions() ([]*domain.Transaction, error) {
	return s.financialTransactionRepo.GetAllTransactions()
}
