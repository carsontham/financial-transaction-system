package handlers

import (
	"encoding/json"
	"errors"
	"financial-transaction-system/app/adapter/http/handlers/viewmodel"
	"financial-transaction-system/app/adapter/http/rest"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/usecase"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CreateNewTransaction(service *usecase.Service, v *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("In handler layer - Creating new transaction")

		var transactionReq viewmodel.TransactionRequest
		if err := json.NewDecoder(req.Body).Decode(&transactionReq); err != nil {
			rest.BadRequest(w)
			return
		}

		if err := v.Struct(transactionReq); err != nil {
			if ve, ok := err.(validator.ValidationErrors); ok {
				rest.UnprocessableEntity(w, ve)
			} else {
				rest.InternalServerError(w)
			}
			return
		}
		// Assumes that idempotency key is generated on client-side and sent in request Header
		idempotencyKey := req.Header.Get(rest.HeaderIdempotencyKey)
		if idempotencyKey == "" {
			idempotencyKey = "hello"
		}

		txn, err := service.GetTransactionByIdempotencyKey(idempotencyKey)
		if err != nil {
			if !errors.Is(err, domain.ErrNotFound) {
				rest.InternalServerError(w)
				return
			}
		}
		// transaction already performed, returns 200 and idempotent result
		if txn != nil {
			rest.StatusOK(w, txn)
			return
		}

		transaction, _ := domain.TransactionViewModelToDomainModel(&transactionReq)

		if err := service.PerformTransaction(transaction); err != nil {
			if errors.Is(err, domain.ErrInsufficientBalance) {
				rest.StatusConflict(w) //409
				return
			}
			rest.InternalServerError(w)
			return
		}
		rest.StatusCreated(w)
		return
	}
}
