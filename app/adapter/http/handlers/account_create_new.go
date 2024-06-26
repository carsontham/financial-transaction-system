package handlers

import (
	"encoding/json"
	"financial-transaction-system/app/adapter/http/handlers/viewmodel"
	"financial-transaction-system/app/adapter/http/rest"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/usecase"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func CreateNewAccount(service usecase.FinancialTransactionService, v *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var accountReq viewmodel.AccountRequest
		if err := json.NewDecoder(req.Body).Decode(&accountReq); err != nil {
			rest.BadRequest(w, err)
			return
		}

		if err := v.Struct(accountReq); err != nil {
			if ve, ok := err.(validator.ValidationErrors); ok {
				rest.UnprocessableEntity(w, ve)
			} else {
				rest.InternalServerError(w)
			}
			return
		}

		account, _ := domain.AccountViewModelToDomainModel(&accountReq)

		// a form of idempotency check
		accExist := usecase.CheckIfAccountExist(service, account.AccountID)
		if accExist {
			rest.StatusOK(w, "account created")
			return
		}

		err := service.CreateNewAccount(account)
		if err != nil {
			log.Println(err)
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, "account created")
	}
}
