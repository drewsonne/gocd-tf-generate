package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestPipeline(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*getPipelineStatusCommand(),
		*pausePipelineCommand(),
		*unpausePipelineCommand(),
		*releasePipelineLockCommand(),
		*getPipelineCommand(),
		*getPipelineHistoryCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Pipelines")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
