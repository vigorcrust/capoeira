package main

import (
	"os"
	"sort"

	"github.com/vigorcrust/capoeira/util"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = util.Name
	app.Version = util.Version
	app.Author = "vigorcrust"
	app.Email = "ddmartens@gmail.com"
	app.Usage = "is an application which lets you tunnel through restrictive proxies."

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	//Sort the flags and commands so that they are sorted when you print the help
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
