package cli

import (
	"context"
	"errors"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	EncryptCommandName  = "encrypt"
	EncryptCommandUsage = "Encrypt a value for use in GoCD configurations"
)

// EncryptAction gets a list of agents and return them.
func EncryptAction(c *cli.Context) error {
	value := c.String("value")
	if value == "" {
		return handleOutput(nil, nil, "Encrypt", errors.New("'--value' is missing"))
	}

	encryptedValue, r, err := cliAgent(c).Encryption.Encrypt(context.Background(), value)
	if err != nil {
		return handleOutput(nil, r, "Encrypt", err)
	}
	return handleOutput(encryptedValue, r, "Encrypt", err)
}

// EncryptCommand checks a template-name is provided and that the response is a 2xx response.
func EncryptCommand() *cli.Command {
	return &cli.Command{
		Name:     EncryptCommandName,
		Usage:    EncryptCommandUsage,
		Action:   EncryptAction,
		Category: "Encryption",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "value"},
		},
	}
}
