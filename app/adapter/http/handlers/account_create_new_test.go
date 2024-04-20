package handlers_test

import (
	"bytes"
	"errors"
	"financial-transaction-system/app/adapter/http/handlers"
	"financial-transaction-system/app/adapter/http/rest"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/tests/servicetest"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNewAccount(t *testing.T) {
	setUp := func(t *testing.T) (
		serviceMock *servicetest.MockFinancialTransactionService,
		c *http.Client,
		url string,
	) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		serviceMock = servicetest.NewMockFinancialTransactionService(ctrl)
		v, _ := rest.NewCustomValidator()

		router := chi.NewRouter()
		router.Post("/accounts", handlers.CreateNewAccount(serviceMock, v))
		s := httptest.NewServer(router)
		c = s.Client()
		url = s.URL + "/accounts"
		return
	}

	t.Run("it should successfully create a new account, return 200", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
				"account_id": 123,
				"initial_balance": "100.23344"
			}`
		stubAccount := &domain.Account{
			AccountID: 123,
			Balance:   100.23344,
		}
		serviceMock.EXPECT().GetAccountById(int64(123)).Times(1).Return(nil, domain.ErrNotFound)
		serviceMock.EXPECT().CreateNewAccount(stubAccount).Times(1).Return(nil)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("it should return status 200 as account id already exist", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
				"account_id": 123,
				"initial_balance": "100.23344"
			}`
		stubAccount := &domain.Account{
			AccountID: 123,
			Balance:   100.23344,
		}
		serviceMock.EXPECT().GetAccountById(int64(123)).Times(1).Return(stubAccount, nil)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("it should fail due to valid_balance validation tag - balance is not float, return 422", func(t *testing.T) {
		_, client, url := setUp(t)
		reqBody := `
			{
				"account_id": 123,
				"initial_balance": "100.23344_abc"
			}`

		expectRespBody :=
			`{ 	"status_code": 422, 
				"error": {"AccountRequest.Balance":"Key: 'AccountRequest.Balance' Error:Field validation for 'Balance' failed on the 'valid_balance' tag"}
			}`

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
			var b bytes.Buffer
			_, _ = b.ReadFrom(resp.Body)
			assert.JSONEq(t, expectRespBody, b.String())
		}
	})

	t.Run("it should fail due to valid_amount validation tag - amount is negative, return 422", func(t *testing.T) {
		_, client, url := setUp(t)
		reqBody := `
			{
				"account_id": 123,
				"initial_balance": "-100.23344"
			}`

		expectRespBody :=
			`{ 	"status_code": 422, 
				"error": {"AccountRequest.Balance":"Key: 'AccountRequest.Balance' Error:Field validation for 'Balance' failed on the 'valid_amount' tag"}
			}`

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
			var b bytes.Buffer
			_, _ = b.ReadFrom(resp.Body)
			assert.JSONEq(t, expectRespBody, b.String())
		}
	})

	t.Run("it should fail due to bad request, return 400 ", func(t *testing.T) {
		_, client, url := setUp(t)
		reqBody := `
			{
				"really_bad_request": 123,
			}`

		expectRespBody :=
			`{ 	"status_code": 400, 
				"error": "invalid character '}' looking for beginning of object key string"
			}`

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			var b bytes.Buffer
			_, _ = b.ReadFrom(resp.Body)
			assert.JSONEq(t, expectRespBody, b.String())
		}
	})

	t.Run("it should fail due to internal server error, return 500 ", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
				"account_id": 123,
				"initial_balance": "100.23344"
			}`
		stubAccount := &domain.Account{
			AccountID: 123,
			Balance:   100.23344,
		}
		stubError := errors.New("unexpected internal error")
		serviceMock.EXPECT().GetAccountById(int64(123)).Times(1).Return(nil, domain.ErrNotFound)
		serviceMock.EXPECT().CreateNewAccount(stubAccount).Times(1).Return(stubError)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		}
	})
}
