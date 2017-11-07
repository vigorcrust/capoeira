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
		Usage:  "is the server which is located in the target network (probably internet).",
		Action: command.CmdServer,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "port, p",
				Usage:  "listening port like 8443",
				EnvVar: "CAPOEIRA_PORT",
			},
		},
	},
	{
		Name:   "client",
		Usage:  "is the client which is located in the source network (probably corpnetwork).",
		Action: command.CmdClient,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "url, u",
				Usage:  "server endpoint url. Only https is allowed like https://server:8443",
				EnvVar: "CAPOEIRA_URL",
			},
		},
	},
}

//CommandNotFound is triggered when the command is not found
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
