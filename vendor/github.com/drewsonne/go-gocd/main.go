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

	app := cli.NewApp()
	app.Name = GoCDUtilityName
	app.Usage = GoCDUtilityUsageInstructions
	app.Version = Version
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		*gocli.ConfigureCommand(),
		*gocli.ListAgentsCommand(),
		*gocli.ListPipelineTemplatesCommand(),
		*gocli.GetAgentCommand(),
		*gocli.GetPipelineTemplateCommand(),
		*gocli.CreatePipelineTemplateCommand(),
		*gocli.UpdateAgentCommand(),
		*gocli.UpdateAgentsCommand(),
		*gocli.UpdatePipelineConfigCommand(),
		*gocli.UpdatePipelineTemplateCommand(),
		*gocli.DeleteAgentCommand(),
		*gocli.DeleteAgentsCommand(),
		*gocli.DeletePipelineTemplateCommand(),
		*gocli.DeletePipelineConfigCommand(),
		*gocli.ListPipelineGroupsCommand(),
		*gocli.GetPipelineHistoryCommand(),
		*gocli.GetPipelineCommand(),
		*gocli.CreatePipelineConfigCommand(),
		*gocli.GenerateJSONSchemaCommand(),
		*gocli.GetPipelineStatusCommand(),
		*gocli.PausePipelineCommand(),
		*gocli.UnpausePipelineCommand(),
		*gocli.ReleasePipelineLockCommand(),
		*gocli.GetConfigurationCommand(),
		*gocli.EncryptCommand(),
		*gocli.GetVersionCommand(),
		*gocli.ListPluginsCommand(),
		*gocli.GetPluginCommand(),
		*gocli.ListScheduledJobsCommand(),
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "server", EnvVar: gocd.EnvVarServer},
		cli.StringFlag{Name: "username", EnvVar: gocd.EnvVarUsername},
		cli.StringFlag{Name: "password", EnvVar: gocd.EnvVarPassword},
		cli.BoolFlag{Name: "ssl_check", EnvVar: gocd.EnvVarSkipSsl},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
