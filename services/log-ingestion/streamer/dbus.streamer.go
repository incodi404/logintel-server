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

func DbusStreamer(ctx context.Context, logCh <-chan models.DbusUnitRecord, errCh chan<- error) {

	// taking log process time from env
	logProcessTime := os.Getenv("LOG_PROCESS_TIME")
	if logProcessTime == "" || !utils.IsNumber(logProcessTime) {
		logProcessTime = "5"
	}

	logProcessTimeNum, _ := strconv.Atoi(logProcessTime)

	ticker := time.NewTicker(time.Duration(logProcessTimeNum) * time.Second)
	defer ticker.Stop()

	esInstance := elasticsearch.Get()
	_, js, err := natsjs.Get()
	if err != nil || js == nil {
		errCh <- err
		return
	}

	// batch of logs
	var logBatch []models.DbusUnitRecord

	for {
		select {
		case <-ctx.Done():
			if len(logBatch) > 0 {
				// process remaining logs before exit
				fmt.Println("[INFO] Flushing remaining logs on shutdown: ", len(logBatch))
			}
			return
		case log, ok := <-logCh:
			if !ok {
				errCh <- fmt.Errorf("[ERROR] Error occured in DbusStreamer log fetch")
			}

			logBatch = append(logBatch, log)

			jsonData, err := json.Marshal(log)
			if err == nil {
				ack, err := js.Publish(ctx, natsjs.DbusDS.Subject, jsonData)
				if err != nil {
					errCh <- err
				}

				fmt.Printf("[DBUS] Stored in %s, sequence %d\n", ack.Stream, ack.Sequence)
			}
		case <-ticker.C:
			if len(logBatch) == 0 {
				continue
			}

			fmt.Println("[INFO] Running log batch after 5 sec: ", len(logBatch))
			// fmt.Println(logBatch)
			if err := elasticsearch.BulkIndex(
				ctx, esInstance, elasticsearch.DbusIndexConfig.IndexName, logBatch,
			); err != nil {
				errCh <- fmt.Errorf("[ES ERROR] Error occured while uploading DBUS logs: %w", err)
			}
			fmt.Println("[INFO] Logs uploaded!")
			logBatch = nil
		}
	}
}
