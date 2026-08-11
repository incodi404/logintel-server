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
		Durable: "exec",
	}

	ExecveDS = DataStreamConfig{
		Name:    "LOGEXECVE",
		Subject: "log.execve.>",
		Durable: "execve",
	}

	DbusDS = DataStreamConfig{
		Name:    "LOGDBUS",
		Subject: "log.dbus.>",
		Durable: "dbus",
	}

	Connect4DS = DataStreamConfig{
		Name:    "LOGCONNECT4",
		Subject: "log.connect4.>",
		Durable: "connect4",
	}

	Bind4DS = DataStreamConfig{
		Name:    "LOGBIND4",
		Subject: "log.bind4.>",
		Durable: "bind4",
	}

	ISSSDS = DataStreamConfig{
		Name:    "LOGISSS",
		Subject: "log.isss.>",
		Durable: "isss",
	}

	FanotifyDS = DataStreamConfig{
		Name:    "LOGFANOTIFY",
		Subject: "log.fanotify.>",
		Durable: "fanotify",
	}
)
