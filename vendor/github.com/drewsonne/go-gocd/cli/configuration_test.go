package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestConfiguration(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*getConfigurationCommand(),
		*getVersionCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Configuration")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
