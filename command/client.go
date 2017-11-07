package command

import (
	"fmt"

	"github.com/urfave/cli"
)

//CmdClient is the main entrypoint for the client
func CmdClient(c *cli.Context) error {
	fmt.Println(c.String("url"))

	return nil
}
