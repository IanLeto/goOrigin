package cron

import (
	"context"
	"goOrigin/pkg"
	"time"
)

//
func TickerAdapter(ctx context.Context, t pkg.Task, interval int64, parallel bool, killChan chan struct{}) {
	var ticker = time.NewTicker(time.Duration(interval))
	for {
		select {
		case <-ticker.C:
			if parallel {
				go func() {
					_ = t.Run(ctx)
				}()
			} else {
				_ = t.Run(ctx)
			}
		case <-killChan:
			return
		case <-ctx.Done():
			return
		}
	}
}
