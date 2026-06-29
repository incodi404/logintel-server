package streamer

import (
	"context"
	"fmt"
	"log-ingestion/elasticsearch"
	"log-ingestion/models"
	"log-ingestion/utils"
	"os"
	"strconv"
	"time"
)

func ExecStreamer(ctx context.Context, logCh <-chan models.ExecRecord, errCh chan<- error) {
	// taking log process time from env
	logProcessTime := os.Getenv("LOG_PROCESS_TIME")
	if logProcessTime == "" || !utils.IsNumber(logProcessTime) {
		logProcessTime = "5"
	}

	logProcessTimeNum, _ := strconv.Atoi(logProcessTime)

	ticker := time.NewTicker(time.Duration(logProcessTimeNum) * time.Second)
	defer ticker.Stop()

	esInstance := elasticsearch.Get()

	var logBatch []models.ExecRecord

	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-logCh:
			if !ok {
				errCh <- fmt.Errorf("[ERROR] Error occured in ExecStreamer log fetch")
			}

			logBatch = append(logBatch, log)

		case <-ticker.C:
			if len(logBatch) == 0 {
				continue
			}

			fmt.Println("[INFO] Running log batch after 5 sec: ", len(logBatch))
			if err := elasticsearch.BulkIndex(
				ctx, esInstance, elasticsearch.ExecIndexConfig.IndexName, logBatch,
			); err != nil {
				errCh <- fmt.Errorf("[ES ERROR] Error occured while uploading EXEC logs: %w", err)
			}
			fmt.Println("[INFO] Logs uploaded!")
			logBatch = nil
		}
	}
}
