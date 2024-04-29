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
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateNewTransaction(t *testing.T) {
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
		router.Post("/transactions", handlers.CreateNewTransaction(serviceMock, v))
		s := httptest.NewServer(router)
		c = s.Client()
		url = s.URL + "/transactions"
		return
	}

	t.Run("it should successfully perform transfer transactions, return 200", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
            "source_account_id": 123,
            "destination_account_id": 456,
            "amount": "100.12345"
			}`
		stubTransaction := &domain.Transaction{
			SourceAccountID:      123,
			DestinationAccountID: 456,
			Amount:               100.12345,
		}
		stubTransaction.IdempotencyKey = gomock.Any().String()

		serviceMock.EXPECT().GetTransactionByIdempotencyKey(gomock.Any()).Times(1).Return(nil, domain.ErrNotFound)
		serviceMock.EXPECT().PerformTransaction(gomock.Any()).Times(1).Return(nil)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			require.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("it should return idempotent response for transactions performed, return 200", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
            "source_account_id": 123,
            "destination_account_id": 456,
            "amount": "100.12345"
			}`
		stubTransaction := &domain.Transaction{
			SourceAccountID:      123,
			DestinationAccountID: 456,
			Amount:               100.12345,
		}
		stubTransaction.IdempotencyKey = "temp-key" +
			strconv.FormatInt(stubTransaction.SourceAccountID, 10) +
			strconv.FormatInt(stubTransaction.SourceAccountID, 10)

		serviceMock.EXPECT().GetTransactionByIdempotencyKey(gomock.Any()).Times(1).Return(stubTransaction, nil)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			require.Equal(t, http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("it should fail due to bad request payload, return 400", func(t *testing.T) {
		_, client, url := setUp(t)
		reqBody := `
			{
            "really_bad_payload": "123",
			}`
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("it should fail due to validation errors, return 422", func(t *testing.T) {
		_, client, url := setUp(t)
		reqBody := `
			{
            "source_account_id": 123,
            "destination_account_id": 456,
            "amount": "-100.12345"
			}`

		expectRespBody :=
			`{ 	"status_code": 422, 
				"error": {"TransactionRequest.Amount":"Key: 'TransactionRequest.Amount' Error:Field validation for 'Amount' failed on the 'valid_amount' tag"}
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

	t.Run("it should fail due to unexpected internal server errors, return 500", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
            "source_account_id": 123,
            "destination_account_id": 456,
            "amount": "100.12345"
			}`
		stubError := errors.New("unexpected internal error")
		serviceMock.EXPECT().GetTransactionByIdempotencyKey(gomock.Any()).Times(1).Return(nil, stubError)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		}
	})

	t.Run("it should fail due to one account is not valid, return 404", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
            "source_account_id": 123,
            "destination_account_id": 999, 
            "amount": "100.12345"
			}`

		//stubTransaction := &domain.Transaction{
		//	SourceAccountID:      123,
		//	DestinationAccountID: 999,
		//	Amount:               100.12345,
		//}
		//stubTransaction.IdempotencyKey = "specific-uuid-string"

		serviceMock.EXPECT().GetTransactionByIdempotencyKey(gomock.Any()).Times(1).Return(nil, domain.ErrNotFound)
		serviceMock.EXPECT().PerformTransaction(gomock.Any()).Times(1).Return(domain.ErrNotFound)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("it should fail due to insufficient balance, return 409", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
            "source_account_id": 123,
            "destination_account_id": 999, 
            "amount": "100.12345"
			}`

		//stubTransaction := &domain.Transaction{
		//	SourceAccountID:      123,
		//	DestinationAccountID: 999,
		//	Amount:               100.12345,
		//}
		//stubTransaction.IdempotencyKey = "temp-key" +
		//	strconv.FormatInt(stubTransaction.SourceAccountID, 10) +
		//	strconv.FormatInt(stubTransaction.SourceAccountID, 10) + time.Now().Format("2006-01-02 15:04:05")

		serviceMock.EXPECT().GetTransactionByIdempotencyKey(gomock.Any()).Times(1).Return(nil, domain.ErrNotFound)
		serviceMock.EXPECT().PerformTransaction(gomock.Any()).Times(1).Return(domain.ErrInsufficientBalance)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusConflict, resp.StatusCode)
		}
	})

	t.Run("it should fail due to internal server error during perform transaction, return 500", func(t *testing.T) {
		serviceMock, client, url := setUp(t)
		reqBody := `
			{
            "source_account_id": 123,
            "destination_account_id": 999, 
            "amount": "100.12345"
			}`

		//stubTransaction := &domain.Transaction{
		//	SourceAccountID:      123,
		//	DestinationAccountID: 999,
		//	Amount:               100.12345,
		//}
		//stubTransaction.IdempotencyKey = "temp-key" +
		//	strconv.FormatInt(stubTransaction.SourceAccountID, 10) +
		//	strconv.FormatInt(stubTransaction.SourceAccountID, 10) + time.Now().Format("2006-01-02 15:04:05")

		stubError := errors.New("unexpected internal error")
		serviceMock.EXPECT().GetTransactionByIdempotencyKey(gomock.Any()).Times(1).Return(nil, domain.ErrNotFound)
		serviceMock.EXPECT().PerformTransaction(gomock.Any()).Times(1).Return(stubError)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		resp, err := client.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		}
	})
}
