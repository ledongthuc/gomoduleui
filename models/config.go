package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Setup  ConfigSetup                 `yaml:"setup"`
	Module map[ModuleName]ConfigModule `yaml:"module"`
}

type ConfigSetup struct {
	GoPrivate      string `yaml:"go_private"`
	CustomCommands string `yaml:"custom_commands"`
	Hidden         bool   `yaml:"hidden"`
}

type ModuleName string

type ConfigModule struct {
	DefaultReplaces []ConfigModuleDefaultReplace       `yaml:"default_replaces"`
	DefaultUpdates  []ConfigModuleDefaultUpdateVersion `yaml:"default_updates"`
	Hidden          bool                               `yaml:"hidden"`
}

type ConfigModuleDefaultReplace string
type ConfigModuleDefaultUpdateVersion string

func ParseConfigurationsFromFile(filePath string) (Config, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	return ParseConfigurations(content)
}

func ParseConfigurations(content []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(content, &config)
	return config, err
}
