package util

import (
	"fmt"

	"github.com/urfave/cli"
)

//PrintErrorAndUsage prints a message and then shows the
//usage help text of context which has been passed.
func PrintErrorAndUsage(m string, ctx *cli.Context) {
	fmt.Fprintf(ctx.App.Writer, "%s\n\n", m)
	cli.ShowCommandHelp(ctx, ctx.Command.FullName())
}
