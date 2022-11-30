package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/libreoffice-docker/unoserver-rest-api/api"
	"github.com/libreoffice-docker/unoserver-rest-api/deport"
	"github.com/libreoffice-docker/unoserver-rest-api/unoconvert"
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
			EnvVar: "UNOCONVERT_EXECUTABLE_PATH",
		},
		cli.DurationFlag{
			Name:  "unoconvert-timeout",
			Value: 0 * time.Minute,
			Usage: "Set the unoconvert run timeout",
		},
	}
	app.Authors = []cli.Author{
		{
			Name:  "libreoffice-docker",
			Email: "https://github.com/libreoffice-docker/unoserver-rest-api",
		},
	}
	app.Action = mainAction

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func mainAction(c *cli.Context) {
	// Create temporary working directory
	deport.MkdirTemp()

	// Cleanup temporary working directory after finished
	defer deport.CleanTemp()

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
