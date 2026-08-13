package logpipelines

import (
	"context"
	"fmt"
	"rule-engine/models"
	natsjs "rule-engine/nats-js"
	"sync"
)

func DbusProcesser(ctx context.Context, errCh chan<- error, dataCh <-chan models.DbusUnitRecord) {
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-dataCh:
			if !ok {
				errCh <- fmt.Errorf("[ERROR] Error fetching log")
				return
			}

			fmt.Println("[RECEIVED] Dbus Log Received!")
			fmt.Println(log)
		}
	}
}

func DbusPipeline(ctx context.Context, errCh chan<- error, wg *sync.WaitGroup) {
	dataCh := make(chan models.DbusUnitRecord, 1000)

	_, js, err := natsjs.Get()
	if err != nil {
		errCh <- err
		return
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(dataCh)
		// consumer
		natsjs.LogConsumer(ctx, js, &natsjs.DbusDS, errCh, dataCh)
	}()

	go func() {
		defer wg.Done()
		// processor
		DbusProcesser(ctx, errCh, dataCh)
	}()
}
