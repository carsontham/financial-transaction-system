package handlers

import (
	"financial-transaction-system/app/usecase"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateNewTransaction(t *testing.T) {
	type args struct {
		service usecase.FinancialTransactionService
		v       *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CreateNewTransaction(tt.args.service, tt.args.v), "CreateNewTransaction(%v, %v)", tt.args.service, tt.args.v)
		})
	}
}
