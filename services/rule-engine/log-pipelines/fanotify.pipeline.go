package logpipelines

import (
	"context"
	"fmt"
	"rule-engine/models"
	natsjs "rule-engine/nats-js"
	"sync"
)

func FanotifyProcesser(ctx context.Context, errCh chan<- error, dataCh <-chan models.FanotifyRecord) {
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-dataCh:
			if !ok {
				errCh <- fmt.Errorf("[ERROR] Error fetching log")
				return
			}

			fmt.Println("[RECEIVED] Fanotify Log Received!")
			fmt.Println(log)
		}
	}
}

func FanotifyPipeline(ctx context.Context, errCh chan<- error, wg *sync.WaitGroup) {
	dataCh := make(chan models.FanotifyRecord, 1000)

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
		natsjs.LogConsumer(ctx, js, &natsjs.FanotifyDS, errCh, dataCh)
	}()

	go func() {
		defer wg.Done()
		// processor
		FanotifyProcesser(ctx, errCh, dataCh)
	}()
}
