package streamer

import (
	"context"
	"fmt"
	"log-ingestion/models"
)

func ExecStreamer(ctx context.Context, logCh <-chan models.ExecRecord, errCh chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-logCh:
			if !ok {
				errCh <- fmt.Errorf("[ERROR] Error occured in ExecStreamer log fetch")
			}

			fmt.Println("[INFO] Exec log received: ", log)
		}
	}
}
