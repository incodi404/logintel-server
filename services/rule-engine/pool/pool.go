package pool

import (
	"context"
	"fmt"
	logpipelines "rule-engine/log-pipelines"
	"sync"
)

func StartWorkerPool(ctx context.Context) {
	// err channels
	fanotifyErrCh := make(chan error, 500)
	dbusErrCh := make(chan error, 500)

	// wg
	var wg *sync.WaitGroup

	// pipelines
	logpipelines.FanotifyPipeline(ctx, fanotifyErrCh, wg)
	logpipelines.DbusPipeline(ctx, dbusErrCh, wg)

	// error pipelines
	LogErrorPipeline(ctx, fanotifyErrCh, "FANOTIFY", wg)
	LogErrorPipeline(ctx, dbusErrCh, "DBUS", wg)

	// close err channels
	go func() {
		wg.Wait()
		close(fanotifyErrCh)
		close(dbusErrCh)
	}()
}

func LogErrorPipeline(ctx context.Context, errCh <-chan error, subject string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err, _ := <-errCh:
				fmt.Printf("[ERROR] Error in error channel :: [%s] :: %s", subject, err)
			}
		}
	}()
}
