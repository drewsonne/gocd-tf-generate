package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestJSON(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*generateJSONSchemaCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Schema")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
