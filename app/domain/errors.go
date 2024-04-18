package domain

import "errors"

var (
	ErrNotFound           = errors.New("account not found")
	ErrParseStringToFloat = errors.New("error parsing balance string to balance float")
)
