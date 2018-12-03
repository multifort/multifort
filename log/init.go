package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
	Initialize log system
*/
func InitLogger(logFile string, logLevel string) *zap.Logger {

	hook := lumberjack.Logger{
		Filename:   logFile, //log file
		MaxSize:    128,     //megabytes (mb)
		MaxBackups: 30,      //30 backups
		MaxAge:     7,       //stay 7 days
		Compress:   true,    //compress
	}

	ws := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch logLevel{
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		ws,
		level,
	)

	logger := zap.New(core)
	logger.Info("Default logger init success")

	return logger
}
