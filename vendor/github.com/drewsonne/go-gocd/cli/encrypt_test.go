package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestEncrypt(t *testing.T) {
	for _, envCmd := range []cli.Command{
		*encryptCommand(),
	} {
		assert.Equal(t, envCmd.Category, "Encryption")
		assert.NotEmpty(t, envCmd.Name)
		assert.NotEmpty(t, envCmd.Usage)
	}
}
