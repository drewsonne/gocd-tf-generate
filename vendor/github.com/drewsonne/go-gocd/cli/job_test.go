package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestJob(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*listScheduledJobsCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Jobs")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
