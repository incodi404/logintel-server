package streamer

import (
	"context"
	"fmt"
	"log-ingestion/models"
)

func DbusStreamer(ctx context.Context, logCh <-chan models.DbusUnitRecord, errCh chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-logCh:
			if !ok {
				errCh <- fmt.Errorf("[ERROR] Error occured in DbusStreamer log fetch")
			}

			fmt.Println("[INFO] Dbus log received: ", log)
		}
	}
}
