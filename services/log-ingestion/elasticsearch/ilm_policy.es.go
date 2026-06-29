package elasticsearch

import (
	"context"
	"fmt"
	"net/http"
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

func (c *Client) IsILMPolicyExist(ctx context.Context, policyName string) (bool, error) {
	res, err := c.ES.ILM.GetLifecycle(
		c.ES.ILM.GetLifecycle.WithContext(ctx),
		c.ES.ILM.GetLifecycle.WithPolicy(policyName),
	)
	if err != nil {
		return false, fmt.Errorf("[ES ERROR] Error occured while checking existed ILM policy: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return false, fmt.Errorf("[ES ERROR] Failed to get ILM policy")
	}

	if res.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return true, nil
}

func (c *Client) CreateILMPolicy(ctx context.Context, policyName string) error {

	isExisted, _ := c.IsILMPolicyExist(ctx, policyName)
	if isExisted {
		return nil
	}

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
