package logger

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	log = logger.Sugar()
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags)
	err := log.Sync()
	if err != nil {
		return
	}
}

func Error(msg string, tags ...zap.Field) {
	log.Error(msg, tags)
	err := log.Sync()
	if err != nil {
		return
	}
}
