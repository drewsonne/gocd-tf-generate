package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestPipelineGroup(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*listPipelineGroupsCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Pipeline Groups")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
