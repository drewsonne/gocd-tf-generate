package cli

import (
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// List of command name and descriptions
const (
	ConfigureCommandName  = "configure"
	ConfigureCommandUsage = "Generate configuration file ~/.gocd.conf"
)

func configureAction(c *cli.Context) error {
	var s string
	var err error
	if s, err = generateConfigFile(); err != nil {
		return err
	}

	if err = ioutil.WriteFile(gocd.ConfigFilePath(),
		[]byte(s), 0644); err != nil {
		return err
	}

	return nil
}

// Build a default template
func generateConfigFile() (string, error) {
	cfg := gocd.Configuration{}

	qs := []*survey.Question{
		{
			Name:     "gocd_server",
			Prompt:   &survey.Input{Message: "GoCD Server (should contain '/go/' suffix)"},
			Validate: survey.Required,
		},
		{
			Name:   "username",
			Prompt: &survey.Input{Message: "Client Username"},
		},
		{
			Name:   "password",
			Prompt: &survey.Password{Message: "Client Password"},
		},
		{
			Name:   "ssl_check",
			Prompt: &survey.Confirm{Message: "Validate SSL Certiicate?"},
		},
	}

	a := struct {
		GoCDServer string `survey:"gocd_server"`
		Username   string
		Password   string
		SslCheck   bool `survey:"ssl_check"`
	}{}

	survey.Ask(qs, &a)

	cfg.Server = a.GoCDServer
	cfg.Username = a.Username
	cfg.Password = a.Password
	cfg.SslCheck = a.SslCheck

	s, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

// ConfigureCommand handles the interaction between the cli flags and the action handler for configure
func ConfigureCommand() *cli.Command {
	return &cli.Command{
		Name:   ConfigureCommandName,
		Usage:  ConfigureCommandUsage,
		Action: configureAction,
	}
}
