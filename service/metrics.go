package service

import "github.com/go-kit/kit/metrics"

var (
	UpGauge               metrics.Gauge
	OutOfBalance          metrics.Gauge
	GasExceedsAllowance   metrics.Gauge
	FailingGauge          metrics.Gauge
	CompletedTransfers    metrics.Counter
	TokenBankBalanceGauge metrics.Gauge
	CoinBankBalanceGauge  metrics.Gauge
)
