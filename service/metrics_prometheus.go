// +build prometheus

package service

import (
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "bridge"
)

func init() {
	UpGauge = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "up",
		Help:      "Is set when application is running.",
	}, []string{})
	OutOfBalance = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "out_of_balance",
		Help:      "Is set when one of the banks will run out of coins.",
	}, []string{})
	GasExceedsAllowance = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "exceeds_allowance",
		Help:      "Is set when any blockchain transaction failed with gas exceeds allowance warning.",
	}, []string{})
	FailingGauge = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "failing",
		Help:      "Is set if application fails with unexpected error.",
	}, []string{})
}
