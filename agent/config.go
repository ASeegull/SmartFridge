package agent

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Config includes config
type Config struct {
	Port      string `yaml:"port"`
	Host      string `yaml:"host"`
	Websocket string `yaml:"websocket"`
	RestURI   string `yaml:"restURI"`
}

const (
	configPath = "./config.yaml"
)

//ReadConfig reads config from file
func ReadConfig() (*Config, error) {
	config := &Config{}
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
