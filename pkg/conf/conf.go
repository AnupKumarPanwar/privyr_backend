package conf

import (
	"time"

	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

type App struct {
	Name        string
	LogSavePath string
	LogSaveName string
}

var AppConf = &App{}

type Server struct {
	RunMode      string
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConf = &Server{}

type Database struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	SslMode  string
}

var DatabaseConf = &Database{}

var cfg *ini.File

func Setup() {

	var err error

	cfg, err = ini.Load("conf.ini")
	if err != nil {
		zap.L().Fatal("Failed to parse conf.ini", zap.String("err", err.Error()))
	}

	mapTo("app", AppConf)
	mapTo("server", ServerConf)
	mapTo("database", DatabaseConf)

	ServerConf.ReadTimeout = ServerConf.ReadTimeout * time.Second
	ServerConf.WriteTimeout = ServerConf.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		zap.L().Fatal("Failed to map configurations")
	}
}
