package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/vigorcrust/capoeira/command"
)

//GlobalFlags holds all global flags
var GlobalFlags = []cli.Flag{}

//Commands holds all subcommands
var Commands = []cli.Command{
	{
		Name:   "server",
		Usage:  "",
		Action: command.CmdServer,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "port",
				Value: "8443",
				Usage: "listening port",
			},
		},
	},
	{
		Name:   "client",
		Usage:  "",
		Action: command.CmdClient,
		Flags:  []cli.Flag{},
	},
}

//CommandNotFound is triggered when the command is not found
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
