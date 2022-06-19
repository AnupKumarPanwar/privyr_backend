package cmd

import (
	"bytes"
	"os/exec"
	"privyr/pkg/conf"

	"go.uber.org/zap"
)

func DbDrop() {
	cmd := exec.Command("dropdb", "-h", conf.DatabaseConf.Host, "-U", conf.DatabaseConf.Username, conf.DatabaseConf.DbName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}
	zap.L().Info("Database dropped successfully")
}
