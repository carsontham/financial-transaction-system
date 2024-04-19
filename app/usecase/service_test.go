package usecase

import (
	"errors"
	"financial-transaction-system/app/domain"
	repositorytest "financial-transaction-system/tests"
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

func TestService_CheckIfAccountExist(t *testing.T) {
	setup := func(t *testing.T) (
		repoMock *repositorytest.MockFinancialTransactionRepository,
		service *Service,
	) {
		ctrl := gomock.NewController(t)
		repoMock = repositorytest.NewMockFinancialTransactionRepository(ctrl)
		service = NewService(repoMock)
		return
	}

	t.Run("it should return false as account does not exists", func(t *testing.T) {
		repo, service := setup(t)
		repo.EXPECT().GetAccountByID(int64(123)).Return(nil, domain.ErrNotFound).Times(1)
		expectFalse := service.CheckIfAccountExist(int64(123))
		assert.Equal(t, false, expectFalse)
	})

	t.Run("it should return true as account exists", func(t *testing.T) {
		repo, service := setup(t)
		stubAccount := &domain.Account{
			AccountID: 123,
			Balance:   100.12345,
		}
		repo.EXPECT().GetAccountByID(int64(123)).Return(stubAccount, nil).Times(1)
		expectTrue := service.CheckIfAccountExist(int64(123))
		assert.Equal(t, true, expectTrue)
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
		actualTxn, err := service.GetTransaction(key)
		assert.NoError(t, err)
		assert.Equal(t, stubTxn, actualTxn)
	})

	t.Run("it should fail and return an error", func(t *testing.T) {
		repo, service := setup(t)
		key := "testKey"
		stubError := errors.New("error")
		repo.EXPECT().GetTransactionByIdempotencyKey(key).Return(nil, stubError).Times(1)
		_, err := service.GetTransaction(key)
		require.ErrorIs(t, err, stubError)
	})
}

func TestService_PerformTransaction(t *testing.T) {

}

func TestService_GetAllTransactions(t *testing.T) {

}