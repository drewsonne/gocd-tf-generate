package cli

import (
	"context"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListPluginsCommandName  = "list-plugins"
	ListPluginsCommandUsage = "List all the Plugins"
	GetPluginCommandName    = "get-plugin"
	GetPluginCommandUsage   = "Get a Plugin"
)

// GetPluginAction retrieves a single plugin by name
func GetPluginAction(c *cli.Context) error {
	pgs, r, err := cliAgent(c).Plugins.Get(context.Background(), c.String("name"))
	if err != nil {
		return handleOutput(nil, r, "GetPlugin", err)
	}

	return handleOutput(pgs, r, "ListPipelineTemplates", err)
}

// ListPluginsAction retrieves all plugin configurations
func ListPluginsAction(c *cli.Context) error {
	pgs, r, err := cliAgent(c).Plugins.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListPlugins", err)
	}

	return handleOutput(pgs, r, "ListPlugins", err)
}

// GetPluginCommand Describes the cli interface for the GetPluginAction
func GetPluginCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPluginCommandName,
		Usage:    GetPluginCommandUsage,
		Category: "Plugins",
		Action:   GetPluginAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
	}
}

// ListPluginsCommand Describes the cli interface for the ListPluginsCommand
func ListPluginsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPluginsCommandName,
		Usage:    ListPluginsCommandUsage,
		Category: "Plugins",
		Action:   ListPluginsAction,
	}
}
