package handlers

import (
	"encoding/json"
	"errors"
	"financial-transaction-system/app/adapter/http/handlers/viewmodel"
	"financial-transaction-system/app/adapter/http/rest"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/usecase"
	"fmt"
	"net/http"
)

func CreateNewAccount(service *usecase.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("In handler layer - Creating new accounts")

		var account viewmodel.AccountRequest
		if err := json.NewDecoder(req.Body).Decode(&account); err != nil {
			rest.BadRequest(w)
			return
		}
		// a form of idempotency check
		isFirstReq := service.IsFirstRequest(account.AccountID)

		if !isFirstReq {
			rest.StatusOK(w, nil)
			return
		}

		err := service.CreateNewAccount(&account)
		if err != nil {
			if errors.Is(err, domain.ErrParseStringToFloat) {
				rest.UnprocessableEntity(w)
				return
			}
			rest.InternalServerError(w)
			return
		}

		rest.StatusCreated(w)
	}
}
