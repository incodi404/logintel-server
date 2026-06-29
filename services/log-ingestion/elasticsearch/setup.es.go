package elasticsearch

import (
	"context"
)

func (c *Client) SetupES(ctx context.Context) error {

	// creatiing policy
	err := c.CreateILMPolicy(ctx, ILM_POLICY_NAME)
	if err != nil {
		return err
	}

	// creating template indexes
	type TemplateConfig struct {
		IndexPatten  string
		IndexName    string
		TemplateName string
		Mapping      string
	}

	templates := []TemplateConfig{
		{ExecIndexConfig.IndexPattern, ExecIndexConfig.IndexName, ExecIndexConfig.TemplateName, ExecMapping},
		{ExecveIndexConfig.IndexPattern, ExecveIndexConfig.IndexName, ExecveIndexConfig.TemplateName, ExecveMapping},
		{Connect4IndexConfig.IndexPattern, Connect4IndexConfig.IndexName, Connect4IndexConfig.TemplateName, Connect4Mapping},
		{Bind4IndexConfig.IndexPattern, Bind4IndexConfig.IndexName, Bind4IndexConfig.TemplateName, Bind4Mapping},
		{ISSSIndexConfig.IndexPattern, ISSSIndexConfig.IndexName, ISSSIndexConfig.TemplateName, ISSSMapping},
		{FanotifyIndexConfig.IndexPattern, FanotifyIndexConfig.IndexName, FanotifyIndexConfig.TemplateName, FanotifyMapping},
		{DbusIndexConfig.IndexPattern, DbusIndexConfig.IndexName, DbusIndexConfig.TemplateName, DbusUnitMapping},
	}

	for _, t := range templates {
		if err := c.CreateIndexTemplate(
			ctx,
			t.TemplateName,
			t.IndexPatten,
			ILM_POLICY_NAME,
			t.Mapping,
		); err != nil {
			return err
		}

		// creating data streams
		if err = c.CreateDataStream(
			ctx, t.IndexName,
		); err != nil {
			return err
		}
	}

	return nil
}
