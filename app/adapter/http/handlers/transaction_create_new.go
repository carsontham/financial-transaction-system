package handlers

import (
	"fmt"
	"net/http"
)

func CreateNewTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("In handler layer - Creating new transaction")

	}
}
