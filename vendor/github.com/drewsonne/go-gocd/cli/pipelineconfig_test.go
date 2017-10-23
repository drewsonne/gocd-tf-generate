package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestPipelineConfig(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*createPipelineConfigCommand(),
		*updatePipelineConfigCommand(),
		*deletePipelineConfigCommand(),
		*getPipelineConfigCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Pipeline Configs")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
