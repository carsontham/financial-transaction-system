package handlers

import (
	"financial-transaction-system/app/adapter/http/rest"
	"financial-transaction-system/app/domain"
	"financial-transaction-system/app/usecase"
	"fmt"
	"net/http"
)

func GetAllTransactions(service usecase.FinancialTransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("In handler layer - retrieving all transactions")

		transactions, err := service.GetAllTransactions()
		if err != nil {
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, domain.TransactionDomainArrayToViewArray(transactions))
	}
}
