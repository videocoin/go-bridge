package client

import (
	"github.com/ethereum/go-ethereum/common"
)

type TransferInfo struct {
	LocalTxHash   common.Hash
	ForeignTxHash common.Hash
}
