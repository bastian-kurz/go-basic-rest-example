package logger

import (
	"github.com/bastian-kurz/basic-rest-example/internal/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"sync"
)

var once sync.Once

var loggerInstance *zap.Logger

func Log() *zap.Logger {
	once.Do(func() {
		env := util.GetStringOrDefault("APP_ENV", "develop")
		switch env {
		case "develop":
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			loggerInstance = logger
		case "production":
			config := zap.NewProductionEncoderConfig()
			config.EncodeTime = zapcore.ISO8601TimeEncoder
			fileEncoder := zapcore.NewJSONEncoder(config)
			logFile, _ := os.OpenFile("prod.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			writer := zapcore.AddSync(logFile)
			defaultLogLevel := zapcore.ErrorLevel
			core := zapcore.NewTee(
				zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			)
			loggerInstance = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		default:
			loggerInstance = zap.NewNop()
		}
	})

	return loggerInstance
}
