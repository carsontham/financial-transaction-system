package handlers_test

import (
	"errors"
	"financial-transaction-system/app/adapter/http/handlers"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/tests/servicetest"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllTransactions(t *testing.T) {
	setUp := func(t *testing.T) (
		serviceMock *servicetest.MockFinancialTransactionService,
		c *http.Client,
		url string,
	) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		serviceMock = servicetest.NewMockFinancialTransactionService(ctrl)

		router := chi.NewRouter()
		router.Get("/transactions", handlers.GetAllTransactions(serviceMock))
		s := httptest.NewServer(router)
		c = s.Client()
		url = s.URL + "/transactions"
		return
	}

	t.Run("it should successfully retrieve all transactions, return 200", func(t *testing.T) {
		serviceMock, client, url := setUp(t)

		stubTransactions := []*domain.Transaction{
			{
				TransactionID:        001,
				SourceAccountID:      123,
				DestinationAccountID: 321,
				Amount:               50.5000,
				IdempotencyKey:       "stubKey",
			},
			{
				TransactionID:        002,
				SourceAccountID:      123,
				DestinationAccountID: 321,
				Amount:               50.5000,
				IdempotencyKey:       "stubKey",
			},
			{
				TransactionID:        003,
				SourceAccountID:      123,
				DestinationAccountID: 321,
				Amount:               50.5000,
				IdempotencyKey:       "stubKey",
			},
		}
		serviceMock.EXPECT().GetAllTransactions().Times(1).Return(stubTransactions, nil)

		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			require.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("it should fail due to unexpected internal server error, return 500", func(t *testing.T) {
		serviceMock, client, url := setUp(t)

		stubError := errors.New("unexpected error")
		serviceMock.EXPECT().GetAllTransactions().Times(1).Return(nil, stubError)

		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		}
	})

}
