package command

import (
	"fmt"

	"github.com/urfave/cli"
)

//CmdServer is the main entrypoint for the server
func CmdServer(c *cli.Context) error {
	fmt.Println(c.String("port"))

	return nil
}
