package db

type DBConfig struct {
	Db       string
	Username string
	Password string
	Max_Conn int16
	Min_Conn int16
}

func SetConfig() {}
