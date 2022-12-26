package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Logger() *zap.Logger {
	lc := zap.NewDevelopmentConfig()
	lc.DisableStacktrace = true
	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := lc.Build()
	return logger
}

func Environment() {
	//read env variable
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			Logger().Error("error while loading .env file", zap.Error(err))
		}
	} else {
		Logger().Info("running service without configuration from .env")
	}
}
