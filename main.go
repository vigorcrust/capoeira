package main

import (
	"os"

	"github.com/vigorcrust/capoeira/util"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = util.Name
	app.Version = util.Version
	app.Author = "vigorcrust"
	app.Email = "ddmartens@gmail.com"
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
