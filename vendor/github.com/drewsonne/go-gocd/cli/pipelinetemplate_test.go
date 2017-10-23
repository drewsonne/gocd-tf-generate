package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestPipelineTemplate(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*deletePipelineTemplateCommand(),
		*listPipelineTemplatesCommand(),
		*getPipelineTemplateCommand(),
		*createPipelineTemplateCommand(),
		*updatePipelineTemplateCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Pipeline Templates")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
