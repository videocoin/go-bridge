package service

import (
	"errors"
	"strings"
)

var (
	// ErrBankOutOfBalance raised if bank doesn't have enough balance to cover transaction.
	ErrBankOutOfBalance = errors.New("bank out of balance")
)

const (
	exceedsAllowancesPattern = "gas required exceeds allowance"
)

func IsErrExceedsAllowance(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), exceedsAllowancesPattern)
}
