package service

import (
	"context"
	"time"
)

// PollForever polls indefinitely.
func PollForever(ctx context.Context, period, timeout time.Duration, f func(context.Context)) error {
	timer := time.NewTimer(period)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			ctx, cancel := context.WithTimeout(ctx, timeout)
			f(ctx)
			cancel()
			timer.Reset(period)
		}
	}
}
