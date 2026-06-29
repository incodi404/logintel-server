package elasticsearch

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) IsDataStreamExist(ctx context.Context, streamName string) (bool, error) {
	res, err := c.ES.Indices.GetDataStream(
		c.ES.Indices.GetDataStream.WithContext(ctx),
		c.ES.Indices.GetDataStream.WithName(streamName),
	)
	if err != nil {
		return false, fmt.Errorf("[ES ERROR] Error occured while checking existed data stream: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return false, fmt.Errorf("[ES ERROR] Failed to get data stream")
	}

	if res.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return true, nil
}

func (c *Client) CreateDataStream(ctx context.Context, streamName string) error {
	isExisted, _ := c.IsDataStreamExist(ctx, streamName)
	if isExisted {
		return nil
	}

	res, err := c.ES.Indices.CreateDataStream(
		streamName,
		c.ES.Indices.CreateDataStream.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("[ES ERRORData stream creation has been failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[ES ERROR] Error occured in Data stream: %w", err)
	}

	return nil
}
