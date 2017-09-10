package cli

import (
	"context"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListPipelineGroupsCommandName  = "list-pipeline-groups"
	ListPipelineGroupsCommandUsage = "List Pipeline Groups"
)

// ListPipelineGroupsAction handles the interaction between the cli flags and the action handler for
// list-pipeline-groups
func ListPipelineGroupsAction(c *cli.Context) error {
	pgs, r, err := cliAgent(c).PipelineGroups.List(context.Background(), c.String("group-name"))
	if err != nil {
		return handleOutput(nil, r, "ListPipelineTemplates", err)
	}

	return handleOutput(pgs, r, "ListPipelineTemplates", err)
}

// ListPipelineGroupsCommand handles the interaction between the cli flags and the action handler for
// list-pipeline-groups
func ListPipelineGroupsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPipelineGroupsCommandName,
		Usage:    ListPipelineGroupsCommandUsage,
		Action:   ListPipelineGroupsAction,
		Category: "Pipeline Groups",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "group-name"},
		},
	}
}
