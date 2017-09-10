package cli

import (
	"context"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	GetConfigurationCommandName  = "get-configuration"
	GetConfigurationCommandUsage = "Get GoCD server configuration. This is the cruise-config.xml file. It is exposed in a json format to enable a consistent format. This API is for read-only purposes and not intended as an interface to modify the config."
	GetVersionCommandName        = "get-version"
	GetVersionCommandUsage       = "Return the Version for the GoCD Server"
)

// GetConfigurationAction gets a list of agents and return them.
func GetConfigurationAction(c *cli.Context) error {
	pgs, r, err := cliAgent(c).Configuration.Get(context.Background())
	if err != nil {
		return handleOutput(nil, r, "GetConfiguration", err)
	}

	return handleOutput(pgs, r, "GetConfiguration", err)
}

// GetVersionAction returns version information about GoCD
func GetVersionAction(c *cli.Context) error {
	v, r, err := cliAgent(c).Configuration.GetVersion(context.Background())
	if err != nil {
		return handleOutput(nil, r, "GetVersion", err)
	}

	return handleOutput(v, r, "GetVersion", err)

}

// GetConfigurationCommand handles the interaction between the cli flags and the action handler for delete-agents
func GetConfigurationCommand() *cli.Command {
	return &cli.Command{
		Name:     GetConfigurationCommandName,
		Usage:    GetConfigurationCommandUsage,
		Action:   GetConfigurationAction,
		Category: "Configuration",
	}
}

// GetVersionCommand handles the interaction between the cli flags and the action handler for delete-agents
func GetVersionCommand() *cli.Command {
	return &cli.Command{
		Name:     GetVersionCommandName,
		Usage:    GetVersionCommandUsage,
		Action:   GetVersionAction,
		Category: "Configuration",
	}
}
