package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestPlugin(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*getPluginCommand(),
		*listPluginsCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Plugins")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
