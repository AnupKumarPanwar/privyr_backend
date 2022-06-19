package cmd

import (
	"bytes"
	"os/exec"
	"privyr/pkg/conf"

	"go.uber.org/zap"
)

func DbCreate() {
	cmd := exec.Command("createdb", "-p", conf.DatabaseConf.Port, "-h", conf.DatabaseConf.Host, "-U", conf.DatabaseConf.Username, "-e", conf.DatabaseConf.DbName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}
	zap.L().Info("Database created successfully")
}
