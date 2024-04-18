package viewmodel

type Account struct {
	AccountID int64   `json:"account_id" validate:"required"`
	Balance   float64 `json:"balance" validate:"required"`
}
