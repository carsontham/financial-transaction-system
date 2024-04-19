package handlers

import (
	"errors"
	"financial-transaction-system/app/adapter/http/rest"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/usecase"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func GetAccountByID(service usecase.FinancialTransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("In handler layer - Getting Account By ID")

		accountID, _ := strconv.ParseInt(chi.URLParam(req, "account_id"), 10, 64)
		acc, err := service.GetAccountById(accountID)
		if err != nil {
			log.Println(err)
			if errors.Is(err, domain.ErrNotFound) {
				rest.NotFound(w)
				return
			}
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, domain.AccountDomainModelToViewModel(acc))
	}
}
