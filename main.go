package main

import (
	"os"
	"privyr/cmd"
	"privyr/constants"
)

func main() {
	if len(os.Args) != 2 {
		cmd.Serve()
	} else {
		command := os.Args[1]
		switch command {
		case constants.Serve:
			cmd.Serve()
			break
		case constants.DbCreate:
			cmd.DbCreate()
			break
		case constants.DbMigrate:
			cmd.DbMigrate()
			break
		case constants.DbDrop:
			cmd.DbDrop()
			break
		default:
			cmd.Serve()
		}
	}
}
