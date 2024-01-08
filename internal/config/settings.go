package config

import (
	"os"

	"github.com/juju/errors"
	"gopkg.in/yaml.v2"
)

var Settings SettingsRoot

type SettingsRoot struct {
	Database Database `yaml:"database"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

func ParseSettings() error {
	f, err := os.Open(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(yaml.NewDecoder(f).Decode(&Settings))
}
