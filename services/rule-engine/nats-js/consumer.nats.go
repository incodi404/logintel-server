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

	info, err := cons.Info(ctx)
	if err != nil {
		errCh <- fmt.Errorf("[NATS] Error getting message info from consumer: %w", err)
		return
	}

	pending := info.NumPending
	if pending == 0 {
		fmt.Println("[INFO] Nothing is pending")
		return
	}

	msgs, err := cons.Fetch(int(pending))
	if err != nil {
		errCh <- fmt.Errorf("[NATS] Error getting messages from consumer: %w", err)
		return
	}

	for msg := range msgs.Messages() {
		var log T

		meta, err := msg.Metadata()
		if err != nil {
			errCh <- fmt.Errorf("[NATS] Error getting message metadata: %w", err)
			continue
		}

		if err = json.Unmarshal(msg.Data(), &log); err != nil {
			errCh <- fmt.Errorf("[NATS] Error unmarshal data: %w", err)
			continue
		}

		// Acknowledge so the server advances the consumer past this message.
		if err := msg.Ack(); err != nil {
			errCh <- fmt.Errorf("[NATS] Error ack data: %w", err)
			continue
		}

		fmt.Printf("[NATS] Stream seq: %d, Consumer seq: %d\n", meta.Sequence.Stream, meta.Sequence.Consumer)
		dataCh <- log // sending to channel
	}
}
