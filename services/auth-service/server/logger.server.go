package server

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitializeLogger() {
	// disable color
	gin.DisableConsoleColor()

	gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "../logs/gin.log",
		MaxSize:    100,
		MaxAge:     5,
		MaxBackups: 3,
		LocalTime:  true,
	}
}
