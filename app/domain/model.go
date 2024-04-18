package domain

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
