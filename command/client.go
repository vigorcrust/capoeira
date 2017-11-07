package command

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vigorcrust/capoeira/util"
)

//CmdClient is the main entrypoint for the client
func CmdClient(c *cli.Context) error {
	if c.NArg() == 0 {
		util.PrintErrorAndUsage("", c)
		return nil
	}
	fmt.Println(c.String("url"))
	return nil
}
