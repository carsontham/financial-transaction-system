package handlers_test

import (
	"errors"
	"financial-transaction-system/app/adapter/http/handlers"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/tests/servicetest"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountByID(t *testing.T) {
	setUp := func(t *testing.T) (
		serviceMock *servicetest.MockFinancialTransactionService,
		c *http.Client,
		url string,
	) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		serviceMock = servicetest.NewMockFinancialTransactionService(ctrl)

		router := chi.NewRouter()
		router.Get("/accounts/{account_id:\\d+}", handlers.GetAccountByID(serviceMock))
		s := httptest.NewServer(router)
		c = s.Client()
		url = s.URL + "/accounts/%d"
		return
	}

	t.Run("it should successfully get an account by id, return 200", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		stubAccountID := int64(123)
		serviceMock.EXPECT().GetAccountById(stubAccountID).Times(1).Return(new(domain.Account), nil)

		req, _ := http.NewRequest("GET", fmt.Sprintf(url, stubAccountID), nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			require.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("it should fail as account does not exist, return 404", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		stubAccountID := int64(123)
		serviceMock.EXPECT().GetAccountById(stubAccountID).Times(1).Return(nil, domain.ErrNotFound)

		req, _ := http.NewRequest("GET", fmt.Sprintf(url, stubAccountID), nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			require.Equal(t, http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("it should fail due to unexpected internal errors, return 500", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		stubAccountID := int64(123)
		stubUnexpectedError := errors.New("unexpected internal server error")
		serviceMock.EXPECT().GetAccountById(stubAccountID).Times(1).Return(nil, stubUnexpectedError)

		req, _ := http.NewRequest("GET", fmt.Sprintf(url, stubAccountID), nil)
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		}
	})
}
