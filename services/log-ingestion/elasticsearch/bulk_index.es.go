package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func BulkIndex[T any](ctx context.Context, c *Client, dataStream string, events []T) error {
	var buf strings.Builder // building string by adding multiple bytes

	for _, event := range events {
		meta := `{"create": {}}` + "\n" // starting line for every log
		buf.WriteString(meta)

		// json conversion
		jsonData, err := json.Marshal(event)
		if err != nil {
			return err
		}

		buf.Write(jsonData)
		buf.WriteString("\n")
	}

	res, err := c.ES.Bulk(
		strings.NewReader(buf.String()),
		c.ES.Bulk.WithContext(ctx),
		c.ES.Bulk.WithIndex(dataStream),
	)
	if err != nil {
		return fmt.Errorf("[ES ERROR] Bulk log upload has been failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[ES ERROR] Error occured in bulk log upload: %w", err)
	}

	return nil
}
