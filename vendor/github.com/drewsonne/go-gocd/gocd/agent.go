package gocd

import (
	"context"
	"fmt"
)

// AgentsService describes Actions which can be performed on agents
type AgentsService service

// AgentsResponse describes the structure of the API response when listing collections of agent object.
type AgentsResponse struct {
	Links    *HALLinks `json:"_links,omitempty"`
	Embedded *struct {
		Agents []*Agent `json:"agents"`
	} `json:"_embedded,omitempty"`
}

// Agent represents agent in GoCD.
type Agent struct {
	UUID             string        `json:"uuid,omitempty"`
	Hostname         string        `json:"hostname,omitempty"`
	ElasticAgentID   string        `json:"elastic_agent_id,omitempty"`
	ElasticPluginID  string        `json:"elastic_plugin_id,omitempty"`
	IPAddress        string        `json:"ip_address,omitempty"`
	Sandbox          string        `json:"sandbox,omitempty"`
	OperatingSystem  string        `json:"operating_system,omitempty"`
	FreeSpace        int           `json:"free_space,omitempty"`
	AgentConfigState string        `json:"agent_config_state,omitempty"`
	AgentState       string        `json:"agent_state,omitempty"`
	Resources        []string      `json:"resources,omitempty"`
	Environments     []string      `json:"environments,omitempty"`
	BuildState       string        `json:"build_state,omitempty"`
	BuildDetails     *BuildDetails `json:"build_details,omitempty"`
	Links            *HALLinks     `json:"_links,omitempty,omitempty"`
	client           *Client
}

// AgentBulkUpdate describes the structure for the PUT payload when updating multiple agents
type AgentBulkUpdate struct {
	Uuids            []string                   `json:"uuids"`
	Operations       *AgentBulkOperationsUpdate `json:"operations,omitempty"`
	AgentConfigState string                     `json:"agent_config_state,omitempty"`
}

// AgentBulkOperationsUpdate describes the structure for a single Operation in AgentBulkUpdate the PUT payload when
// updating multiple agents
type AgentBulkOperationsUpdate struct {
	Environments *AgentBulkOperationUpdate `json:"environments,omitempty"`
	Resources    *AgentBulkOperationUpdate `json:"resources,omitempty"`
}

// AgentBulkOperationUpdate describes an action to be performed on an Environment or Resource during an agent update.
type AgentBulkOperationUpdate struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

// BuildDetails describes the builds being performed on this agent.
type BuildDetails struct {
	Links    *HALLinks `json:"_links"`
	Pipeline string    `json:"pipeline"`
	Stage    string    `json:"stage"`
	Job      string    `json:"job"`
}

// List will retrieve all agents, their status, and metadata from the GoCD Server.
func (s *AgentsService) List(ctx context.Context) ([]*Agent, *APIResponse, error) {
	r := AgentsResponse{}
	_, resp, err := s.client.getAction(ctx, &APIClientRequest{
		Path:         "agents",
		ResponseBody: &r,
		APIVersion:   apiV4,
	})

	for _, agent := range r.Embedded.Agents {
		agent.client = s.client
	}

	return r.Embedded.Agents, resp, err
}

// Get will retrieve a single agent based on the provided UUID.
func (s *AgentsService) Get(ctx context.Context, uuid string) (*Agent, *APIResponse, error) {
	return s.handleAgentRequest(ctx, "GET", uuid, nil)
}

// Update will modify the configuration for an existing agents.
func (s *AgentsService) Update(ctx context.Context, uuid string, agent *Agent) (*Agent, *APIResponse, error) {
	return s.handleAgentRequest(ctx, "PATCH", uuid, agent)
}

// Delete will remove an existing agent. Note: The agent must be disabled, and not currently building to be deleted.
func (s *AgentsService) Delete(ctx context.Context, uuid string) (string, *APIResponse, error) {
	return s.client.deleteAction(ctx, "agents/"+uuid, apiV4)
}

// BulkUpdate will change the configuration for multiple agents in a single request.
func (s *AgentsService) BulkUpdate(ctx context.Context, agents AgentBulkUpdate) (string, *APIResponse, error) {
	a := StringResponse{}
	_, resp, err := s.client.patchAction(ctx, &APIClientRequest{
		Path:         "agents",
		APIVersion:   apiV4,
		ResponseBody: &a,
		RequestBody:  agents,
	})
	return a.Message, resp, err
}

// JobRunHistory will return a list of Jobs run on this agent.
func (s *AgentsService) JobRunHistory(ctx context.Context, uuid string) ([]*Job, *APIResponse, error) {
	a := JobRunHistoryResponse{}
	_, resp, err := s.client.getAction(ctx, &APIClientRequest{
		Path:         fmt.Sprintf("agents/%s/job_run_history", uuid),
		APIVersion:   apiV4,
		ResponseBody: &a,
	})
	return a.Jobs, resp, err
}

// handleAgentRequest handles the flow to perform an HTTP action on an agent resource.
func (s *AgentsService) handleAgentRequest(ctx context.Context, action string, uuid string, agent *Agent) (*Agent, *APIResponse, error) {
	a := Agent{client: s.client}
	_, resp, err := s.client.httpAction(ctx, &APIClientRequest{
		Method:       action,
		Path:         "agents/" + uuid,
		APIVersion:   apiV4,
		RequestBody:  agent,
		ResponseBody: &a,
	})

	return &a, resp, err
}
