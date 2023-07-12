package main

import (
	"log"
	"os"

	"github.com/libreofficedocker/unoserver-rest-api/server"
	"github.com/urfave/cli"
)

var Version = "unstable"

func init() {
	log.SetPrefix("unoserver-rest-api ")
}

func main() {
	app := cli.NewApp()
	app.Name = "unoserver-rest-api"
	app.Version = Version
	app.Usage = "The simple REST API for unoserver and unoconvert"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "0.0.0.0:2004",
			Usage: "The addr used by the unoserver api server",
		},
	}
	app.Authors = []cli.Author{
		{
			Name:  "libreofficedocker",
			Email: "https://github.com/libreofficedocker/unoserver-rest-api",
		},
	}
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	s := server.Default()

	return s.Run()
}
