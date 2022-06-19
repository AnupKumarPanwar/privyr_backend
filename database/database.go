package database

import (
	"fmt"
	"privyr/pkg/conf"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", conf.DatabaseConf.Host, conf.DatabaseConf.Username, conf.DatabaseConf.Password, conf.DatabaseConf.DbName, conf.DatabaseConf.Port, conf.DatabaseConf.SslMode)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		zap.L().Fatal("models.Setup", zap.String("err", err.Error()))
	}
}
