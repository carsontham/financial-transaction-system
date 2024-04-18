package domain

import "financial-transaction-system/app/adapter/http/handlers/viewmodel"

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

func AccountDomainModelToViewModel(acc *Account) *viewmodel.Account {
	return &viewmodel.Account{
		AccountID: acc.AccountID,
		Balance:   acc.Balance,
	}
}
