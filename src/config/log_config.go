package config

import (
	log "github.com/sirupsen/logrus"
)

const (
	INFO  = "INFO"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
)

func NewLogConfig(config *Config) {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetReportCaller(true)
	level := config.LogLevel.Level
	switch level {
	case INFO:
		log.SetLevel(log.InfoLevel)
	case ERROR:
		log.SetLevel(log.ErrorLevel)
	case DEBUG:
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.FatalLevel)
	}
}
