package viewmodel

type AccountResponse struct {
	AccountID int64  `json:"account_id"`
	Balance   string `json:"balance"`
}

type AccountRequest struct {
	AccountID int64  `json:"account_id" validate:"required"`
	Balance   string `json:"initial_balance" validate:"required,valid_balance,valid_amount"`
}

type TransactionRequest struct {
	SourceAccountID      int64 `json:"source_account_id"`
	DestinationAccountID int64 `json:"destination_account_id"`
	Amount               int64 `json:"amount"`
}
