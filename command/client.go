package command

import (
	"log"

	"github.com/urfave/cli"
	"github.com/vigorcrust/capoeira/util"
)

//CmdClient is the main entrypoint for the client
func CmdClient(c *cli.Context) error {
	log.SetPrefix("[" + util.Name + "-client] ")
	if c.NumFlags() < 2 {
		util.PrintErrorAndUsage("", c)
		return nil
	}

	if c.GlobalBool("verbose") {
		log.Println(c.String("url"))
	}

	return nil
}
