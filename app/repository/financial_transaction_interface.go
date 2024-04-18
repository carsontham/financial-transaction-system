package repository

// FinancialTransactionRepository is the dependency to be injected into components that require access to the DB
type FinancialTransactionRepository interface {
	GetAccountByID() error
	CreateNewAccount() error
	CreateNewTransaction() error
}
