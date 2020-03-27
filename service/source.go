package service

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

var _ AddressesSource = (StaticSource)(nil)

type StaticSource []common.Address

func (s StaticSource) All(context.Context) ([]common.Address, error) {
	return s, nil
}
