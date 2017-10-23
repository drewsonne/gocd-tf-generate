package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestAgent(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*listAgentsCommand(),
		*getAgentCommand(),
		*updateAgentCommand(),
		*deleteAgentCommand(),
		*updateAgentsCommand(),
		*deleteAgentsCommand(),
	} {
		assert.Equal(t, envCmd.Category, agentCategory)
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
