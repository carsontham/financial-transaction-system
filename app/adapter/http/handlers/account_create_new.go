package handlers

import (
	"encoding/json"
	"financial-transaction-system/app/adapter/http/handlers/viewmodel"
	"financial-transaction-system/app/adapter/http/rest"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/usecase"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CreateNewAccount(service *usecase.Service, v *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var accountReq viewmodel.AccountRequest
		if err := json.NewDecoder(req.Body).Decode(&accountReq); err != nil {
			rest.BadRequest(w)
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
		accExist := service.CheckIfAccountExist(account.AccountID)
		if accExist {
			rest.StatusOK(w, nil)
			return
		}

		err := service.CreateNewAccount(account)
		if err != nil {
			rest.InternalServerError(w)
			return
		}
		rest.StatusCreated(w)
	}
}
