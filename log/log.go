package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	logger = InitLogger(viper.GetString("log.file"), viper.GetString("log.level"))
)

func Debug(msg string, field ...zap.Field)  {
	logger.Debug(msg, field...)
}

func Info(msg string, field ...zap.Field){
	logger.Info(msg, field...)
}

func Warn(msg string, field ...zap.Field){
	logger.Warn(msg, field...)
}

func Error(msg string, field ...zap.Field){
	logger.Error(msg, field...)
}

func Fatal(msg string, field ...zap.Field){
	logger.Fatal(msg, field...)
}