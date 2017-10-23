package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestProperties(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*listPropertiesCommand(),
		*createPropertyCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Properties")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
