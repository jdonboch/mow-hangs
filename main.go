package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("mow-hangs", "example to show mow.cli hang")

	_ = app.String(
		cli.StringOpt{
			Name:   "c configpath",
			Value:  "",
			Desc:   fmt.Sprintf("Specify a custom location where the configuration is located."),
			EnvVar: "FARO_HOME",
		})

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v", err)
	}

}
