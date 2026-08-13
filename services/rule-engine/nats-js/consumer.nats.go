package natsjs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
)

func LogConsumer[T any](ctx context.Context, js jetstream.JetStream, dataStreamConf *DataStreamConfig, errCh chan<- error, dataCh chan<- T) {
	// Bind to the durable consumer created earlier
	cons, err := js.Consumer(ctx, dataStreamConf.Name, dataStreamConf.Durable)
	if err != nil {
		errCh <- fmt.Errorf("[NATS] Error binding with consumer: %w", err)
		return
	}

	fmt.Println("[NATS] Block 1 Consumer [2]")

	consCtx, err := cons.Consume(func(msg jetstream.Msg) {
		var log T
		if err = json.Unmarshal(msg.Data(), &log); err != nil {
			errCh <- fmt.Errorf("[NATS] Error unmarshal data: %w", err)
		}

		fmt.Println("[NATS] Got a new message!")

		dataCh <- log // sending to channel

		// Acknowledge so the server advances the consumer past this message.
		if err := msg.Ack(); err != nil {
			errCh <- fmt.Errorf("[NATS] Error ack data: %w", err)
		}
	})
	if err != nil {
		errCh <- fmt.Errorf("[NATS] Error getting message: %w", err)
		return
	}

	defer consCtx.Stop()

	<-ctx.Done() // Keep this LogConsumer function alive until the parent context is cancelled. It keeps the function alive and waits until the parent ctx is cancelled.
}
