// GoCD CLI tool

package main

import (
	gocli "github.com/drewsonne/go-gocd/cli"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
	"os"
	"sort"
)

// GoCDUtilityName is used in help text to identify the gocd cli util by name
const GoCDUtilityName = "gocd"

// GoCDUtilityUsageInstructions providers user facing support on operation of the gocd cli tool. @TODO Expand this content.
const GoCDUtilityUsageInstructions = "CLI Tool to interact with GoCD server"

// Version for the cli tool
var Version string

func main() {
	buildCli().Run(os.Args)
}

func buildCli() *cli.App {

	app := cli.NewApp()
	app.Name = GoCDUtilityName
	app.Usage = GoCDUtilityUsageInstructions
	app.Version = Version
	app.EnableBashCompletion = true
	app.Commands = gocli.GetCliCommands()

	app.Metadata = map[string]interface{}{}
	app.Metadata["c"] = gocli.NewCliClient

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "profile",
			EnvVar: gocd.EnvVarDefaultProfile,
		},
		cli.StringFlag{
			Name:   "server",
			EnvVar: gocd.EnvVarServer,
		},
		cli.StringFlag{
			Name:   "username",
			EnvVar: gocd.EnvVarUsername,
		},
		cli.StringFlag{
			Name:   "password",
			EnvVar: gocd.EnvVarPassword,
		},
		cli.BoolFlag{
			Name:   "skip_ssl_check",
			EnvVar: gocd.EnvVarSkipSsl,
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
