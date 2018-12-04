package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

func InitLogger() {

	logLevel := viper.GetString("log.level")
	logFile := viper.GetString("log.file")
	var level zapcore.Level

	hook := lumberjack.Logger{
		Filename:   logFile, //log file
		MaxSize:    128,     //megabytes (mb)
		MaxBackups: 30,      //30 backups
		MaxAge:     7,       //stay 7 days
		Compress:   true,    //compress
	}

	ws := zapcore.AddSync(&hook)

	switch logLevel{
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "panic":
		level = zap.PanicLevel
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
	defer logger.Sync()
	sugar = logger.Sugar()
	Info("Default logger init success")
}

func Debug(msg string, args ...interface{})  {
	sugar.Debugf(msg, args)
}

func Info(msg string, args ...interface{}){
	sugar.Infof(msg, args)
}

func Warn(msg string, args ...interface{}){
	sugar.Warnf(msg, args)
}

func Error(msg string, args ...interface{}){
	sugar.Errorf(msg, args)
}

func Panic(msg string, args ...interface{})  {
	sugar.Panicf(msg, args)
}

func Fatal(msg string, args ...interface{}){
	sugar.Fatalf(msg, args)
}