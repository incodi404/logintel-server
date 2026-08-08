package natsjs

type DataStreamConfig struct {
	Name    string
	Subject string
}

var (
	ExecDS = DataStreamConfig{
		Name:    "LOGEXEC",
		Subject: "log.exec.>",
	}

	ExecveDS = DataStreamConfig{
		Name:    "LOGEXECVE",
		Subject: "log.execve.>",
	}

	DbusDS = DataStreamConfig{
		Name:    "LOGDBUS",
		Subject: "log.dbus.>",
	}

	Connect4DS = DataStreamConfig{
		Name:    "LOGCONNECT4",
		Subject: "log.connect4.>",
	}

	Bind4DS = DataStreamConfig{
		Name:    "LOGBIND4",
		Subject: "log.bind4.>",
	}

	ISSSDS = DataStreamConfig{
		Name:    "LOGISSS",
		Subject: "log.isss.>",
	}

	FanotifyDS = DataStreamConfig{
		Name:    "LOGFANOTIFY",
		Subject: "log.fanotify.>",
	}
)
