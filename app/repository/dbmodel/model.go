package dbmodel

import "financial-transaction-system/app/domain"

type Account struct {
	AccountID int64   `db:"id"`
	Balance   float64 `db:"title"`
}

type Transaction struct {
	TransactionID        int64   `db:"transaction_id"`
	SourceAccountID      int64   `db:"source_account_id"`
	DestinationAccountID int64   `db:"destination_account_id"`
	Amount               float64 `db:"amount"`
}

func AccountDBModelToDomainModel(acc Account) *domain.Account {
	return &domain.Account{
		AccountID: acc.AccountID,
		Balance:   acc.Balance,
	}
}
