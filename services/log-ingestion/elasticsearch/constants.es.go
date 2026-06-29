package elasticsearch

const ILM_POLICY_NAME = "logintel-log-policy"

type IndexConfig struct {
	TemplateName string
	IndexPattern string
	IndexName    string
}

var (
	DbusIndexConfig = IndexConfig{
		TemplateName: "dbus-log-template",
		IndexPattern: "dbus-log*",
		IndexName:    "dbus-log",
	}

	ExecIndexConfig = IndexConfig{
		TemplateName: "exec-log-template",
		IndexPattern: "exec-log*",
		IndexName:    "exec-log",
	}

	ExecveIndexConfig = IndexConfig{
		TemplateName: "execve-log-template",
		IndexPattern: "execve-log*",
		IndexName:    "execve-log",
	}

	Connect4IndexConfig = IndexConfig{
		TemplateName: "connect4-log-template",
		IndexPattern: "connect4-log*",
		IndexName:    "connect4-log",
	}

	Bind4IndexConfig = IndexConfig{
		TemplateName: "bind4-log-template",
		IndexPattern: "bind4-log*",
		IndexName:    "bind4-log",
	}

	ISSSIndexConfig = IndexConfig{
		TemplateName: "isss-log-template",
		IndexPattern: "isss-log*",
		IndexName:    "isss-log",
	}

	FanotifyIndexConfig = IndexConfig{
		TemplateName: "fanotify-log-template",
		IndexPattern: "fanotify-log*",
		IndexName:    "fanotify-log",
	}
)
