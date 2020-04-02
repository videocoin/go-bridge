package tokentonative

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/erc20"
	"github.com/videocoin/go-bridge/service"
)

type Client interface {
	bind.DeployBackend
	BalanceAt(context.Context, common.Address, *big.Int) (*big.Int, error)
}

func NewERC20Access(log *logrus.Entry, erc20 *erc20.ERC20) *ERC20Access {
	return &ERC20Access{
		log:   log,
		erc20: erc20,
	}
}

type ERC20Access struct {
	log   *logrus.Entry
	erc20 *erc20.ERC20
}

func (e *ERC20Access) Transfers(ctx context.Context, to []common.Address, start, end uint64) (rst []service.Transfer, err error) {
	filter, err := e.erc20.FilterTransfer(
		&bind.FilterOpts{
			Start: start,
			End:   &end,
		},
		nil, to,
	)
	if err != nil {
		return nil, err
	}
	defer filter.Close()
	for filter.Next() {
		if filter.Event != nil {
			if (filter.Event.From == common.Address{}) {
				e.log.Debugf("transfer 0x%x to zero address is skipped", filter.Event.Raw.TxHash)
				continue
			}

			e.log.Debugf("found transfer 0x%x. to 0x%x. value %v",
				filter.Event.Raw.TxHash,
				filter.Event.From,
				filter.Event.Value,
			)
			rst = append(rst, service.Transfer{
				Hash:  filter.Event.Raw.TxHash,
				To:    filter.Event.From,
				Value: filter.Event.Value,
			})
		}
	}
	if filter.Error() != nil {
		return nil, err
	}
	return rst, nil
}
