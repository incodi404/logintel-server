package streamer

import (
	"context"
	"fmt"
	"log-ingestion/models"
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

			fmt.Println("[INFO] Dbus log received")
			logBatch = append(logBatch, log)

		case <-ticker.C:
			if len(logBatch) == 0 {
				continue
			}

			fmt.Println("[INFO] Running log batch after 5 sec: ", len(logBatch))
			fmt.Println(logBatch)
			logBatch = nil
		}
	}
}
