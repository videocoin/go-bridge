// +build prometheus

package common

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	promCollectTimeout = 10 * time.Second
)

func BootstrapPrometheus(ctx context.Context, addr string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		ReadTimeout: promCollectTimeout,
		Handler:     mux,
	}
	errors := make(chan error, 1)
	go func() {
		errors <- srv.ListenAndServe()
	}()
	go func() {
		<-ctx.Done()
		_ = srv.Close()
	}()
	return <-errors
}
