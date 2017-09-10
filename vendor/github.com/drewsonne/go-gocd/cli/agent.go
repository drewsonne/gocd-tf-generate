package cli

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListAgentsCommandName    = "list-agents"
	ListAgentsCommandUsage   = "List GoCD build agents."
	GetAgentCommandName      = "get-agent"
	GetAgentCommandUsage     = "Get Agent by UUID"
	UpdateAgentCommandName   = "update-agent"
	UpdateAgentCommandUsage  = "Update Agent"
	DeleteAgentCommandName   = "delete-agent"
	DeleteAgentCommandUsage  = "Delete Agent"
	UpdateAgentsCommandName  = "update-agents"
	UpdateAgentsCommandUsage = "Bulk Update Agents"
	DeleteAgentsCommandName  = "delete-agents"
	DeleteAgentsCommandUsage = "Bulk Delete Agents"
)

// ListAgentsAction gets a list of agents and return them.
func listAgentsAction(c *cli.Context) error {
	agents, r, err := cliAgent(c).Agents.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListAgents", err)
	}
	for _, agent := range agents {
		agent.RemoveLinks()
	}
	return handleOutput(agents, r, "ListAgents", err)
}

// GetAgentAction retrieves a single agent object.
func getAgentAction(c *cli.Context) error {
	agent, r, err := cliAgent(c).Agents.Get(context.Background(), c.String("uuid"))
	if r.HTTP.StatusCode != 404 {
		agent.RemoveLinks()
	}
	return handleOutput(agent, r, "GetAgent", err)
}

// UpdateAgentAction updates a single agent.
func updateAgentAction(c *cli.Context) error {

	if c.String("uuid") == "" {
		return handleOutput(nil, nil, "UpdateAgent", errors.New("'--uuid' is missing"))
	}

	if c.String("config") == "" {
		return handleOutput(nil, nil, "UpdateAgent", errors.New("'--config' is missing"))
	}

	a := gocd.AgentUpdate{}
	b := []byte(c.String("config"))
	if err := json.Unmarshal(b, &a); err != nil {
		return handleOutput(nil, nil, "UpdateAgent", err)
	}

	agent, r, err := cliAgent(c).Agents.Update(context.Background(), c.String("uuid"), a)
	if r.HTTP.StatusCode != 404 {
		agent.RemoveLinks()
	}
	return handleOutput(agent, r, "UpdateAgent", err)
}

// DeleteAgentAction delets an agent. Note: The agent must be disabled.
func deleteAgentAction(c *cli.Context) error {
	if c.String("uuid") == "" {
		return handleOutput(nil, nil, "DeleteAgent", errors.New("'--uuid' is missing"))
	}

	deleteResponse, r, err := cliAgent(c).Agents.Delete(context.Background(), c.String("uuid"))
	if r.HTTP.StatusCode == 406 {
		err = errors.New(deleteResponse)
	}
	return handleOutput(deleteResponse, r, "DeleteAgent", err)
}

// UpdateAgentsAction updates a single agent.
func updateAgentsAction(c *cli.Context) error {

	u := gocd.AgentBulkUpdate{}
	if o := c.String("operations"); o != "" {
		b := []byte(o)
		op := gocd.AgentBulkOperationsUpdate{}
		if err := json.Unmarshal(b, &op); err == nil {
			return handleOutput(nil, nil, "BulkAgentUpdate", err)
		}
		u.Operations = &op
	}

	var uuids []string
	if uuids = c.StringSlice("uuid"); len(uuids) == 0 {
		return handleOutput(nil, nil, "BulkAgentUpdate", errors.New("'--uuid' is missing"))
	}
	u.Uuids = uuids

	if state := c.String("state"); state != "" {
		u.AgentConfigState = c.String("state")
	}

	updateResponse, r, err := cliAgent(c).Agents.BulkUpdate(context.Background(), u)
	if r.HTTP.StatusCode == 406 {
		err = errors.New(updateResponse)
	}
	return handleOutput(updateResponse, r, "BulkAgentUpdate", err)
}

// DeleteAgentsAction must be implemented.
func deleteAgentsAction(c *cli.Context) error {
	return nil
}

// ListAgentsCommand checks a template-name is provided and that the response is a 2xx response.
func ListAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListAgentsCommandName,
		Usage:    ListAgentsCommandUsage,
		Action:   listAgentsAction,
		Category: "Agents",
	}
}

// GetAgentCommand handles the interaction between the cli flags and the action handler for get-agent
func GetAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     GetAgentCommandName,
		Usage:    GetAgentCommandUsage,
		Action:   getAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
		},
	}
}

// UpdateAgentCommand handles the interaction between the cli flags and the action handler for update-agent
func UpdateAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdateAgentCommandName,
		Usage:    UpdateAgentCommandUsage,
		Action:   updateAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
			cli.StringFlag{Name: "config, c", Usage: "JSON encoded config for agent update."},
		},
	}
}

// DeleteAgentCommand handles the interaction between the cli flags and the action handler for delete-agent
func DeleteAgentCommand() *cli.Command {
	return &cli.Command{
		Name:     DeleteAgentCommandName,
		Usage:    DeleteAgentCommandUsage,
		Action:   deleteAgentAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "uuid, u", Usage: "GoCD Agent UUID"},
		},
	}
}

// UpdateAgentsCommand handles the interaction between the cli flags and the action handler for update-agents
func UpdateAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdateAgentsCommandName,
		Usage:    UpdateAgentsCommandUsage,
		Action:   updateAgentsAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringSliceFlag{Name: "uuid", Usage: "GoCD Agent UUIDs"},
			cli.StringFlag{Name: "state", Usage: "Whether agents are enabled or disabled. Allowed values 'Enabled','Disabled'."},
			cli.StringFlag{Name: "operations", Usage: "JSON encoded config for bulk operation updates."},
		},
	}
}

// DeleteAgentsCommand handles the interaction between the cli flags and the action handler for delete-agents
func DeleteAgentsCommand() *cli.Command {
	return &cli.Command{
		Name:     DeleteAgentsCommandName,
		Usage:    DeleteAgentsCommandUsage,
		Action:   deleteAgentsAction,
		Category: "Agents",
		Flags: []cli.Flag{
			cli.StringSliceFlag{Name: "uuid", Usage: "GoCD Agent UUIDs"},
		},
	}
}
