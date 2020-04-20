// +build !prometheus

package service

import "github.com/go-kit/kit/metrics/discard"

func init() {
	UpGauge = discard.NewGauge()
	OutOfBalance = discard.NewGauge()
	GasExceedsAllowance = discard.NewGauge()
	FailingGauge = discard.NewGauge()
}
