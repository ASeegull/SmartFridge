package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	configPath = "./consoleClientConfig.yaml"
)

//ConsoleClientConfig includes server config
type ConsoleClientConfig struct {
	ServerAddress string `yaml:"serverAddress"`
}

//ReadConfig reads config from file
func ReadConfig() (*ConsoleClientConfig, error) {
	config := &ConsoleClientConfig{}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
