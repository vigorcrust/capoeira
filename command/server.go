package command

import (
	"log"

	"github.com/urfave/cli"
	"github.com/vigorcrust/capoeira/util"
)

//CmdServer is the main entrypoint for the server
func CmdServer(c *cli.Context) error {
	log.SetPrefix("[" + util.Name + "-server] ")
	if c.NumFlags() < 2 {
		util.PrintErrorAndUsage("", c)
		return nil
	}

	if c.GlobalBool("verbose") {
		log.Println(c.String("port"))
	}

	return nil
}
