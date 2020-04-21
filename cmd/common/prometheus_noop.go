// +build !prometheus

package common

import (
	"context"
)

func BootstrapPrometheus(ctx context.Context, addr string) error {
	return nil
}
