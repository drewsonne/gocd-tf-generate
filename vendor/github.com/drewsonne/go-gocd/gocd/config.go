package gocd

import (
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
	EnvVarServer   = "GOCD_SERVER"
	EnvVarUsername = "GOCD_USERNAME"
	EnvVarPassword = "GOCD_PASSWORD"
	EnvVarSkipSsl  = "GOCD_SKIP_SSL_CHECK"
)

// LoadConfig loads configurations from yaml at default file location
func LoadConfig() (*Configuration, error) {
	var b []byte
	cfg := &Configuration{}

	p := ConfigFilePath()
	if _, err := os.Stat(p); !os.IsNotExist(err) {
		if b, err = ioutil.ReadFile(p); err != nil {
			return nil, err
		}

		if err = yaml.Unmarshal(b, &cfg); err != nil {
			return nil, err
		}
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

	return cfg, nil
}

// ConfigFilePath specifies the default path to a config file
func ConfigFilePath() string {
	// @TODO Make it work for windows. Maybe...
	usr, _ := user.Current()
	return strings.Replace(ConfigDirectoryPath, "~", usr.HomeDir, 1)
}
