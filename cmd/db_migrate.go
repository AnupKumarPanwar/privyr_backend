package cmd

import (
	"database/sql"
	"fmt"
	"privyr/pkg/conf"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type MigrationLogger struct {
}

func (migrationLogger MigrationLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (migrationLogger MigrationLogger) Verbose() bool {
	return true
}

func DbMigrate() {
	db, err := sql.Open(conf.DatabaseConf.Driver, fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s", conf.DatabaseConf.Username, conf.DatabaseConf.DbName, conf.DatabaseConf.SslMode, conf.DatabaseConf.Password))
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		zap.L().Fatal(err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		conf.DatabaseConf.DbName, driver)

	m.Log = &MigrationLogger{}
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		zap.L().Fatal(err.Error())
	}
	zap.L().Info("Database migrated")
}
