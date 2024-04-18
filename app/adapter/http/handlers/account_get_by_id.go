package handlers

import (
	"fmt"
	"net/http"
)

func GetAccountByID() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("In handler layer - Getting Account By ID")

	}
}
