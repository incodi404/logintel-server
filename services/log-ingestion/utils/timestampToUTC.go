package utils

import "time"

func TimestampToUTC(sec int64, nano int64) string {
	return time.Unix(sec, nano).UTC().Format(time.RFC3339Nano)
}
