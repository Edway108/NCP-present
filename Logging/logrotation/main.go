package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./app.log",
		MaxSize:    5, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	})

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	logger := zap.New(core)

	logger.Info("Application started")
	logger.Error("Something went wrong")
}
