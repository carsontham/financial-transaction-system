package domain

import (
	"financial-transaction-system/app/adapter/http/handlers/viewmodel"
	"log"
	"strconv"
)

type Account struct {
	AccountID int64
	Balance   float64
}

type Transaction struct {
	TransactionID        int64
	SourceAccountID      int64
	DestinationAccountID int64
	Amount               float64
}

func AccountDomainModelToViewModel(acc *Account) *viewmodel.AccountResponse {
	return &viewmodel.AccountResponse{
		AccountID: acc.AccountID,
		Balance:   strconv.FormatFloat(acc.Balance, 'f', -1, 64),
	}
}

func AccountViewModelToDomainModel(acc *viewmodel.AccountRequest) (*Account, error) {

	balance, err := strconv.ParseFloat(acc.Balance, 64)
	if err != nil {
		log.Println("error parsing string balance to float balance")
		return nil, err
	}
	return &Account{AccountID: acc.AccountID, Balance: balance}, nil
}
