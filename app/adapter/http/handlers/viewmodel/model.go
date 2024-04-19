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
	SourceAccountID      int64  `json:"source_account_id" validate:"required"`
	DestinationAccountID int64  `json:"destination_account_id" validate:"required"`
	Amount               string `json:"amount" validate:"required,valid_balance,valid_amount"`
}

type TransactionResponse struct {
	TransactionID        int64  `json:"transaction_id"`
	SourceAccountID      int64  `json:"source_account_id"`
	DestinationAccountID int64  `json:"destination_account_id"`
	Amount               string `json:"amount"`
	IdempotencyKey       string `json:"idempotency_key"`
}
