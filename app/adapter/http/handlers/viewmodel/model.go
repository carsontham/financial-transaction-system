package viewmodel

type AccountResponse struct {
	AccountID int64  `json:"account_id"`
	Balance   string `json:"balance"`
}

type AccountRequest struct {
	AccountID int64  `json:"account_id" validate:"required"`
	Balance   string `json:"initial_balance" validate:"required"`
}
