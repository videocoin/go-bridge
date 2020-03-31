package service

import "errors"

var (
	// ErrOutOfBalance raised if bank doesn't have enough balance to cover transaction.
	ErrBankOutOfBalance = errors.New("bank out of balance")
)
