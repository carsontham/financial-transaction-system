package domain

import "errors"

var (
	ErrNotFound            = errors.New("resource not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
)
