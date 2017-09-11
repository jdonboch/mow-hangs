package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("mow-hangs", "example to show mow.cli hang")
	app.Version("v version", "test version 1.0")

	// Define Global Flags
	var (
		account    = app.StringOpt("account", "", "Option to specify an AWS Account")
		assume     = app.StringOpt("a assume", "", "Option to assume an AWS role")
		awsProfile = app.StringOpt("A awsprofile", "", "Option to supply an aws credentials profile (overrides the file)")
		output     = app.StringOpt("o output", "", "File to log output to")
		profile    = app.StringOpt("p P profile", "", "Option to supply a test configuration file profile")
		region     = app.StringOpt("r R region", "", "Option to supply an aws region (overrides the file)")
		debug      = app.BoolOpt("d D debug", false, "Debug mode on")
		registry   = app.StringOpt("registry", "myapp.com", "Overrides the location of the test WS Registry")
		trace      = app.BoolOpt("t T trace", false, "Trace mode on")
		cfghome    = app.String(
			cli.StringOpt{
				Name:   "c configpath",
				Value:  "",
				Desc:   fmt.Sprintf("Specify a custom location where the configuration is located."),
				EnvVar: "FARO_HOME",
			})
	)

	app.Before = func() {

		// Make sure Go doesn't complain about us not using these
		_ = *assume
		_ = *awsProfile
		_ = *output
		_ = *profile
		_ = *region
		_ = *debug
		_ = *trace
		_ = *account
		_ = *registry
		_ = *registry
		_ = *cfghome
	}

	app.Command("api", "Test command", sampleCmd())

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v", err)
	}

}

// sampleCmd is an example command
func sampleCmd() func(cmd *cli.Cmd) {
	return func(cmd *cli.Cmd) {
		var port = cmd.StringOpt("p port", "8080", "Specifies the port")
		cmd.LongDesc = "API starts the web service\n" +
			"\n" +
			"   Example:\n" +
			"      main api -p 9090"

		cmd.Action = func() {
			fmt.Printf("Hello world! at port %v", *port)
		}
	}
}
