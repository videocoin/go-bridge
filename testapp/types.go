package testapp

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate stringer -type=TransferType
type TransferType int

const (
	Deposit TransferType = iota + 1
	Withdraw
)

//go:generate stringer -type=TransferState
type TransferState int

const (
	Failed TransferState = iota + 1
	Success
)

type Transfer struct {
	ID uint

	Type  TransferType
	State TransferState

	LocalHash   common.Hash
	ForeignHash common.Hash

	// Latency to complete full bridge action from a client perspective.
	Latency time.Duration
}
