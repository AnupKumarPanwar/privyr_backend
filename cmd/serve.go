package cmd

import (
	"fmt"
	"privyr/database"
	"privyr/pkg/conf"
	"privyr/pkg/logger"
	"privyr/pkg/logwriter"
	"privyr/router"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	logger.Setup()
	conf.Setup()
	logwriter.Setup()
}

func Serve() {
	database.Setup()
	router.Setup()

	gin.SetMode(conf.ServerConf.RunMode)

	readTimeout := conf.ServerConf.ReadTimeout
	writeTimeout := conf.ServerConf.WriteTimeout
	endPoint := fmt.Sprintf("%s:%d", conf.ServerConf.Host, conf.ServerConf.Port)
	maxHeaderBytes := 1 << 20

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, router.Router)
	server.BeforeBegin = func(add string) {
		zap.L().Info("Process started", zap.Int("pid", syscall.Getpid()))
	}

	zap.L().Info("Starting http server", zap.String("endpoint", endPoint))

	err := server.ListenAndServe()
	if err != nil {
		zap.L().Fatal(err.Error())

	}
}
