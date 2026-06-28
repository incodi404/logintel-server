package elasticsearch

import (
	"context"
	"fmt"
	"strings"
)

const policy = `
{
	"policy": {
		"phases": {
			"hot": {
				"actions": {
					"rollover": {
						"max_size": "10gb"
					}
				}
			},
			"delete": {
				"min_age": "3h",
				"actions": {
					"delete": {}
				}
			}
		}
	}
}
`

func (c *Client) CreateILMPolicy(ctx context.Context, policyName string) error {
	res, err := c.ES.ILM.PutLifecycle(
		policyName,
		c.ES.ILM.PutLifecycle.WithContext(ctx),
		c.ES.ILM.PutLifecycle.WithBody(strings.NewReader(policy)),
	)
	if err != nil {
		return fmt.Errorf("[ES ERROR] Policy creation has been failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[ES ERROR] Error occured in ILM Policy: %w", err)
	}

	return nil
}
