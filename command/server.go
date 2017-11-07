package command

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vigorcrust/capoeira/util"
)

//CmdServer is the main entrypoint for the server
func CmdServer(c *cli.Context) error {
	if c.NArg() == 0 {
		util.PrintErrorAndUsage("", c)
		return nil
	}
	fmt.Println(c.String("port"))

	return nil
}
