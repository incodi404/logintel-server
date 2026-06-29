package elasticsearch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) IsIndexTemplateExist(ctx context.Context, templateName string) (bool, error) {
	res, err := c.ES.Indices.GetIndexTemplate(
		c.ES.Indices.GetIndexTemplate.WithContext(ctx),
		c.ES.Indices.GetIndexTemplate.WithName(templateName),
	)
	if err != nil {
		return false, fmt.Errorf("[ES ERROR] Error occured while checking existed index template: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return false, fmt.Errorf("[ES ERROR] Failed to get index template")
	}

	if res.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return true, nil
}

func (c *Client) CreateIndexTemplate(
	ctx context.Context, templateName string, pattern string, policyName string, mapping string,
) error {

	isExisted, _ := c.IsIndexTemplateExist(ctx, templateName)
	if isExisted {
		return nil
	}

	body := fmt.Sprintf(`{
		"index_patterns": ["%s"],
		"data_stream": {},
		"template": {
			"settings": {
				"index.lifecycle.name": "%s"
			},
			"mappings": %s
		}
	}`, pattern, policyName, mapping)

	res, err := c.ES.Indices.PutIndexTemplate(
		templateName,
		strings.NewReader(body),
		c.ES.Indices.PutIndexTemplate.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("[ES ERROR] Index template creation has been failed: %w", err)
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)

	if res.IsError() {
		return fmt.Errorf("[ES ERROR] Error occured in index template: %s, %s", res.Status(), string(bodyBytes))
	}

	return nil
}
