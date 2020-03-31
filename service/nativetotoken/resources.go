package nativetotoken

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/videocoin/go-bridge/nativeproxy"
	"github.com/videocoin/go-bridge/service"
)

func NewNativeTransferAccess(log *logrus.Entry, proxy *nativeproxy.NativeProxy) *NativeTransfersAccess {
	return &NativeTransfersAccess{
		log:   log,
		proxy: proxy,
	}
}

type NativeTransfersAccess struct {
	log   *logrus.Entry
	proxy *nativeproxy.NativeProxy
}

func (a *NativeTransfersAccess) Transfers(
	ctx context.Context,
	to []common.Address,
	start, end uint64) (rst []service.Transfer, err error) {
	filter, err := a.proxy.FilterTransferProxied(
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
