package main

import (
	"os"

	"github.com/urfave/cli"
)

var release = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "unoserver-rest-api"
	app.Version = release
	app.Usage = "The simple REST API for unoserver"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "executable",
			Value:  "unoconvert",
			Usage:  "Set the unoconvert executable path.",
			EnvVar: "UNOCONVERT_EXECUTABLE_PATH",
		},
	}
	app.Author = "Socheat Sok"
	app.Email = "socheatsok78@gmail.com"
	app.Action = mainAction

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func mainAction(c *cli.Context) {
	// server.ListenAndServe()
}
