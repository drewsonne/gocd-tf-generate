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
	configureCommandName  = "configure"
	configureCommandUsage = "Generate configuration file ~/.gocd.conf"
)

type generateConfigFunc func() (cfg *gocd.Configuration, err error)
type loadConfigFunc func() (cfgs map[string]*gocd.Configuration, err error)
type writeConfigFunc func(b []byte) (err error)

type configActionRunner struct {
	context *cli.Context

	generateConfig generateConfigFunc
	loadConfigs    loadConfigFunc
	writeConfigs   writeConfigFunc
}

func (car configActionRunner) run() (err error) {
	var cfg *gocd.Configuration
	var profile string
	var b []byte

	if profile = car.context.Parent().String("profile"); profile == "" {
		profile = "default"
	}

	cfgs, err := car.loadConfigs()
	if err != nil {
		return NewCliError("Configure:generate", nil, err)
	}

	if cfg, err = car.generateConfig(); err != nil {
		return NewCliError("Configure:generate", nil, err)
	}

	cfgs[profile] = cfg

	if b, err = yaml.Marshal(cfgs); err != nil {
		return NewCliError("Configure:yaml", nil, err)
	}

	car.writeConfigs(b)

	return nil
}

func writeConfigsToFile(b []byte) (err error) {
	path, err := gocd.ConfigFilePath()
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(path, b, 0644); err != nil {
		return NewCliError("Configure:write", nil, err)
	}
	return nil
}

// Build a default template
func generateConfig() (cfg *gocd.Configuration, err error) {
	cfg = &gocd.Configuration{}
	qs := []*survey.Question{
		{
			Name:     "server",
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
			Name:   "skip_ssl_check",
			Prompt: &survey.Confirm{Message: "Skip SSL certificate validation"},
		},
	}

	err = survey.Ask(qs, cfg)
	return
}

// ConfigureCommand handles the interaction between the cli flags and the action handler for configure
func configureCommand() *cli.Command {
	return &cli.Command{
		Name:  configureCommandName,
		Usage: configureCommandUsage,
		Action: func(c *cli.Context) (err error) {
			return configActionRunner{
				context: c,

				generateConfig: generateConfig,
				loadConfigs:    gocd.LoadConfigFromFile,
				writeConfigs:   writeConfigsToFile,
			}.run()
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "profile"},
		},
	}
}
