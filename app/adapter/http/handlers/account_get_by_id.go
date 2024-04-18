package handlers

import (
	"financial-transaction-system/app/usecase"
	"fmt"
	"net/http"
)

func GetAccountByID(service *usecase.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("In handler layer - Getting Account By ID")

		acc, err := service.GetAccountById(123)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(acc)
	}
}
