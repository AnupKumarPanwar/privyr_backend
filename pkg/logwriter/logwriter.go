package logwriter

import (
	"io"
	"os"
	"privyr/pkg/conf"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

func Setup() {
	logFileName := conf.AppConf.LogSavePath + conf.AppConf.Name + "/" + conf.AppConf.LogSaveName

	lumberjackWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     90,
		Compress:   true,
	})

	consoleWriter := zapcore.Lock(os.Stdout)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), lumberjackWriter),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.Development())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	gin.DefaultWriter = io.MultiWriter(lumberjackWriter, consoleWriter)
}
