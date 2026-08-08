package natsjs

type DataStreamConfig struct {
	Name    string
	Subject string
	Durable string
}

var (
	ExecDS = DataStreamConfig{
		Name:    "LOGEXEC",
		Subject: "log.exec.>",
		Durable: "",
	}

	ExecveDS = DataStreamConfig{
		Name:    "LOGEXECVE",
		Subject: "log.execve.>",
		Durable: "",
	}

	DbusDS = DataStreamConfig{
		Name:    "LOGDBUS",
		Subject: "log.dbus.>",
		Durable: "",
	}

	Connect4DS = DataStreamConfig{
		Name:    "LOGCONNECT4",
		Subject: "log.connect4.>",
		Durable: "",
	}

	Bind4DS = DataStreamConfig{
		Name:    "LOGBIND4",
		Subject: "log.bind4.>",
		Durable: "",
	}

	ISSSDS = DataStreamConfig{
		Name:    "LOGISSS",
		Subject: "log.isss.>",
		Durable: "",
	}

	FanotifyDS = DataStreamConfig{
		Name:    "LOGFANOTIFY",
		Subject: "log.fanotify.>",
		Durable: "",
	}
)
