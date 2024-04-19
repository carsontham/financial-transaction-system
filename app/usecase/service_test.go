package usecase

import (
	"errors"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/tests/repositorytest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_GetAccountById(t *testing.T) {
	setup := func(t *testing.T) (
		repoMock *repositorytest.MockFinancialTransactionRepository,
		service *Service,
	) {
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockFinancialTransactionRepository(ctrl)
		service = NewService(repoMock)
		return
	}

	t.Run("it should return an account by id", func(t *testing.T) {
		repo, service := setup(t)
		stubAccount := &domain.Account{
			AccountID: 123,
			Balance:   100.12345,
		}
		repo.EXPECT().GetAccountByID(int64(123)).Return(stubAccount, nil).Times(1)
		account, err := service.GetAccountById(stubAccount.AccountID)
		assert.NoError(t, err)
		assert.Equal(t, stubAccount, account)
	})
}

func TestService_CreateNewAccount(t *testing.T) {
	setup := func(t *testing.T) (
		repoMock *repositorytest.MockFinancialTransactionRepository,
		service *Service,
	) {
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockFinancialTransactionRepository(ctrl)
		service = NewService(repoMock)
		return
	}

	t.Run("it should create a new account", func(t *testing.T) {
		repo, service := setup(t)
		stubAccount := &domain.Account{
			AccountID: 123,
			Balance:   100.12345,
		}

		repo.EXPECT().CreateNewAccount(stubAccount).Return(nil).Times(1)
		err := service.CreateNewAccount(stubAccount)
		assert.NoError(t, err)
	})

	t.Run("it should return an error during creation", func(t *testing.T) {
		repo, service := setup(t)
		stubAccount := &domain.Account{
			AccountID: 123,
			Balance:   100.12345,
		}

		repo.EXPECT().CreateNewAccount(stubAccount).Return(errors.New("error")).Times(1)
		err := service.CreateNewAccount(stubAccount)
		assert.Equal(t, errors.New("error"), err)
	})

}

func TestService_GetTransaction(t *testing.T) {
	setup := func(t *testing.T) (
		repoMock *repositorytest.MockFinancialTransactionRepository,
		service *Service,
	) {
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockFinancialTransactionRepository(ctrl)
		service = NewService(repoMock)
		return
	}

	t.Run("it should return a transaction object", func(t *testing.T) {
		repo, service := setup(t)
		key := "testKey"
		stubTxn := &domain.Transaction{
			TransactionID:        123,
			SourceAccountID:      888,
			DestinationAccountID: 999,
			Amount:               500.55,
			IdempotencyKey:       key,
		}
		repo.EXPECT().GetTransactionByIdempotencyKey(key).Return(stubTxn, nil).Times(1)
		actualTxn, err := service.GetTransactionByIdempotencyKey(key)
		assert.NoError(t, err)
		assert.Equal(t, stubTxn, actualTxn)
	})

	t.Run("it should fail and return an error", func(t *testing.T) {
		repo, service := setup(t)
		key := "testKey"
		stubError := errors.New("error")
		repo.EXPECT().GetTransactionByIdempotencyKey(key).Return(nil, stubError).Times(1)
		_, err := service.GetTransactionByIdempotencyKey(key)
		require.ErrorIs(t, err, stubError)
	})
}

func TestService_PerformTransaction(t *testing.T) {
	setup := func(t *testing.T) (
		repoMock *repositorytest.MockFinancialTransactionRepository,
		service *Service,
	) {
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockFinancialTransactionRepository(ctrl)
		service = NewService(repoMock)
		return
	}

	t.Run("it should successfully perform a transfer transaction", func(t *testing.T) {
		repo, service := setup(t)
		key := "testKey"
		stubAccount := &domain.Account{
			AccountID: 123,
		}

		stubTxn := &domain.Transaction{
			TransactionID:        123,
			SourceAccountID:      888,
			DestinationAccountID: 999,
			Amount:               500.55,
			IdempotencyKey:       key,
		}

		repo.EXPECT().GetAccountByID(int64(888)).Return(stubAccount, nil).Times(1)
		repo.EXPECT().GetAccountByID(int64(999)).Return(stubAccount, nil).Times(1)
		repo.EXPECT().PerformTransaction(stubTxn).Return(nil).Times(1)
		err := service.PerformTransaction(stubTxn)
		require.NoError(t, err)

	})

	t.Run("it should fail and return an error if any account does not exist", func(t *testing.T) {
		repo, service := setup(t)
		key := "testKey"
		stubError := domain.ErrNotFound

		stubTxn := &domain.Transaction{
			TransactionID:        123,
			SourceAccountID:      888,
			DestinationAccountID: 999,
			Amount:               500.55,
			IdempotencyKey:       key,
		}

		repo.EXPECT().GetAccountByID(int64(888)).Return(nil, stubError).Times(1)
		repo.EXPECT().GetAccountByID(int64(999)).Return(nil, stubError).Times(1)
		err := service.PerformTransaction(stubTxn)
		assert.ErrorIs(t, err, stubError)

	})
}

func TestService_GetAllTransactions(t *testing.T) {
	setup := func(t *testing.T) (
		repoMock *repositorytest.MockFinancialTransactionRepository,
		service *Service,
	) {
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockFinancialTransactionRepository(ctrl)
		service = NewService(repoMock)
		return
	}

	t.Run("it should successfully retrieve all transactions", func(t *testing.T) {
		repo, service := setup(t)
		stubTransactions := []*domain.Transaction{
			{
				TransactionID:        001,
				SourceAccountID:      123,
				DestinationAccountID: 321,
				Amount:               500.55,
				IdempotencyKey:       "testKey",
			},
			{
				TransactionID:        002,
				SourceAccountID:      123,
				DestinationAccountID: 124,
				Amount:               200.55678,
				IdempotencyKey:       "testKey",
			},
		}

		repo.EXPECT().GetAllTransactions().Return(stubTransactions, nil).Times(1)
		transactions, err := service.GetAllTransactions()
		require.NoError(t, err)
		assert.Equal(t, stubTransactions, transactions)
	})
}
