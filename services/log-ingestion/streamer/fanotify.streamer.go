package streamer

import (
	"context"
	"encoding/json"
	"fmt"
	"log-ingestion/elasticsearch"
	"log-ingestion/models"
	natsjs "log-ingestion/nats-js"
	"log-ingestion/utils"
	"os"
	"strconv"
	"time"
)

func FanotifyStreamer(ctx context.Context, logCh <-chan models.FanotifyRecord, errCh chan<- error) {
	// taking log process time from env
	logProcessTime := os.Getenv("LOG_PROCESS_TIME")
	if logProcessTime == "" || !utils.IsNumber(logProcessTime) {
		logProcessTime = "5"
	}

	logProcessTimeNum, _ := strconv.Atoi(logProcessTime)

	ticker := time.NewTicker(time.Duration(logProcessTimeNum) * time.Second)
	defer ticker.Stop()

	// getting instances
	esInstance := elasticsearch.Get() // es
	_, js, err := natsjs.Get()
	if err != nil || js == nil {
		errCh <- err
		return
	}

	var logBatch []models.FanotifyRecord

	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-logCh:
			// This portion runs in everytime the logCh get a log
			if !ok {
				errCh <- fmt.Errorf("[ERROR] Error occured in FanotifyStreamer log fetch")
			}

			// bathcing in an array
			logBatch = append(logBatch, log)

			// sending to JS
			jsonData, err := json.Marshal(log)
			if err == nil {
				ack, err := js.Publish(ctx, natsjs.FanotifyDS.Subject, jsonData)
				if err != nil {
					errCh <- err
				}

				fmt.Printf("[FANOTIFY] Stored in %s, sequence %d\n", ack.Stream, ack.Sequence)
			}

		case <-ticker.C:
			// This portion runs in interval (default: 5 second)
			if len(logBatch) == 0 {
				continue
			}

			fmt.Println("[INFO] Running log batch after 5 sec: ", len(logBatch))
			if err := elasticsearch.BulkIndex(
				ctx, esInstance, elasticsearch.FanotifyIndexConfig.IndexName, logBatch,
			); err != nil {
				errCh <- fmt.Errorf("[ES ERROR] Error occured while uploading FANOTIFY logs: %w", err)
			}
			fmt.Println("[INFO] Logs uploaded!")
			logBatch = nil
		}
	}
}
