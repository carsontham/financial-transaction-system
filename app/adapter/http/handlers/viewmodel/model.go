package viewmodel

type Account struct {
	AccountID int64  `json:"account_id" validate:"required"`
	Balance   string `json:"balance" validate:"required"`
}
