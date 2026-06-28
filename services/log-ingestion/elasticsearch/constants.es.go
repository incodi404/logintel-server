package elasticsearch

const ILM_POLICY_NAME = "logintel-log-policy"

type IndexConfig struct {
	TemplateName string
	IndexPattern string
	IndexName    string
}

var (
	Exec = IndexConfig{
		TemplateName: "exec-log-template",
		IndexPattern: "exec-log-*",
		IndexName:    "exec-log",
	}
)
