package logger

import (
	"go.uber.org/zap"
)

func Setup() {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	zap.ReplaceGlobals(logger)
}
