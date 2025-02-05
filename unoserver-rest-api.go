package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/libreofficedocker/unoserver-rest-api/api"
	"github.com/libreofficedocker/unoserver-rest-api/depot"
	"github.com/libreofficedocker/unoserver-rest-api/unoconvert"
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
			Value: "0.0.0.0:80",
			Usage: "The addr used by the unoserver api server",
		},
		cli.StringFlag{
			Name:   "unoserver-addr",
			Value:  "127.0.0.1:2002",
			Usage:  "The unoserver addr used by the unoconvert",
			EnvVar: "UNOSERVER_ADDR",
		},
		cli.StringFlag{
			Name:   "unoconvert-bin",
			Value:  "unoconvert",
			Usage:  "Set the unoconvert executable path.",
			EnvVar: "UNOCONVERT_BIN",
		},
		cli.DurationFlag{
			Name:   "unoconvert-timeout",
			Value:  0 * time.Minute,
			Usage:  "Set the unoconvert run timeout",
			EnvVar: "UNOCONVERT_TIMEOUT",
		},
	}
	app.Authors = []cli.Author{
		{
			Name:  "libreoffice-docker",
			Email: "https://github.com/libreofficedocker/unoserver-rest-api",
		},
	}
	app.Action = mainAction

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func mainAction(c *cli.Context) {
	// Create temporary working directory
	depot.MkdirTemp()

	// Cleanup temporary working directory after finished
	defer depot.CleanTemp()

	// Configure unoconvert options
	unoAddr := c.String("unoserver-addr")
	host, port, _ := net.SplitHostPort(unoAddr)
	unoconvert.SetInterface(host)
	unoconvert.SetPort(port)
	unoconvert.SetExecutable(c.String("unoconvert-bin"))
	unoconvert.SetContextTimeout(c.Duration("unoconvert-timeout"))

	// Start the API server
	addr := c.String("addr")
	api.ListenAndServe(addr)
}
