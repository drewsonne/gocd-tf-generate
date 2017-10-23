package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestEnvironment(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*listEnvironmentsCommand(),
		*getEnvironmentCommand(),
		*addPipelinesToEnvironmentCommand(),
		*removePipelinesFromEnvironmentCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Environments")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
