package logger

import (
	"sync"

	"go.uber.org/zap"
)


var (
	log *zap.SugaredLogger
	once sync.Once
)

func Init(appEnv string) {
	once.Do(func() {
		var baseLogger *zap.Logger
		var err error

		if appEnv == "production" {
			baseLogger, err = zap.NewProduction()
		} else {
			baseLogger, err = zap.NewDevelopment()
		}

		if err != nil {
			panic("failed to initialize logger: " + err.Error())
		}

		log = baseLogger.Sugar()
	})
}

func Info(msg string, fields ...interface{}) {
	log.Infow(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	log.Errorw(msg, fields...)
}

func Get() *zap.SugaredLogger {
	return log
}