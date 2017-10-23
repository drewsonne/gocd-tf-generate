package gocd

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

// ConfigDirectoryPath is the location where the authentication information is stored
const ConfigDirectoryPath = "~/.gocd.conf"

// Environment variables for configuration.
const (
	EnvVarDefaultProfile = "GOCD_DEFAULT_PROFILE"
	EnvVarServer         = "GOCD_SERVER"
	EnvVarUsername       = "GOCD_USERNAME"
	EnvVarPassword       = "GOCD_PASSWORD"
	EnvVarSkipSsl        = "GOCD_SKIP_SSL_CHECK"
)

// Configuration object used to initialise a gocd lib client to interact with the GoCD server.
type Configuration struct {
	Server       string
	Username     string `yaml:"username,omitempty"`
	Password     string `yaml:"password,omitempty"`
	SkipSslCheck bool   `yaml:"skip_ssl_check,omitempty" survey:"skip_ssl_check"`
}

// LoadConfigByName loads configurations from yaml at default file location
func LoadConfigByName(name string, cfg *Configuration) (err error) {

	cfgs, err := LoadConfigFromFile()
	if err == nil {
		newCfg, hasCfg := cfgs[name]
		if !hasCfg {
			return fmt.Errorf("Could not find configuration profile '%s'", name)
		}

		*cfg = *newCfg
	} else {
		return err
	}

	if server := os.Getenv(EnvVarServer); server != "" {
		cfg.Server = server
	}

	if username := os.Getenv(EnvVarUsername); username != "" {
		cfg.Username = username
	}

	if password := os.Getenv(EnvVarPassword); password != "" {
		cfg.Password = password
	}

	return nil
}

// LoadConfigFromFile on disk and return it as a Config item
func LoadConfigFromFile() (cfgs map[string]*Configuration, err error) {
	var b []byte

	p, err := ConfigFilePath()
	if err != nil {
		return cfgs, err
	}
	if _, err := os.Stat(p); !os.IsNotExist(err) {
		if b, err = ioutil.ReadFile(p); err != nil {
			return nil, err
		}

		if err = yaml.Unmarshal(b, &cfgs); err != nil {
			return nil, err
		}
	}

	return cfgs, nil
}

// ConfigFilePath specifies the default path to a config file
func ConfigFilePath() (string, error) {
	// @TODO Make it work for windows. Maybe...
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := usr.HomeDir
	return strings.Replace(ConfigDirectoryPath, "~", homeDir, 1), nil
}
