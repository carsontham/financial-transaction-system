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
	IdempotencyKey       string  `db:"idempotency_key"`
}

func AccountDBModelToDomainModel(acc Account) *domain.Account {
	return &domain.Account{
		AccountID: acc.AccountID,
		Balance:   acc.Balance,
	}
}

func AccountDomainModelToDBModel(acc *domain.Account) *Account {
	return &Account{
		AccountID: acc.AccountID,
		Balance:   acc.Balance,
	}
}

func TransactionDBModelToDomainModel(txn Transaction) *domain.Transaction {
	return &domain.Transaction{
		TransactionID:        txn.TransactionID,
		SourceAccountID:      txn.SourceAccountID,
		DestinationAccountID: txn.DestinationAccountID,
		Amount:               txn.Amount,
		IdempotencyKey:       txn.IdempotencyKey,
	}
}

func TransactionDomainModelToDBModel(txn *domain.Transaction) *Transaction {
	return &Transaction{
		TransactionID:        txn.TransactionID,
		SourceAccountID:      txn.SourceAccountID,
		DestinationAccountID: txn.DestinationAccountID,
		Amount:               txn.Amount,
		IdempotencyKey:       txn.IdempotencyKey,
	}
}
