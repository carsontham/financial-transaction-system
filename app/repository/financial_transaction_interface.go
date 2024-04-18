package repository

type FinancialTransactionRepository interface {
	GetAccountByID() error
	CreateNewAccount() error
	CreateNewTransaction() error
}
