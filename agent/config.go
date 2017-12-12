package agent

import (
	"io/ioutil"

	"github.com/davecheney/errors"
	"gopkg.in/yaml.v2"
)

//Config includes config
type Config struct {
	Port      string `yaml:"port"`
	Host      string `yaml:"host"`
	Websocket string `yaml:"websocket"`
	RestURI   string `yaml:"restURI"`
}

//ReadConfig reads config from file
func ReadConfig(configPath string) (*Config, error) {
	config := &Config{}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.Annotate(err, "could not read yaml file")
	}

	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return nil, errors.Annotate(err, "could not decode config file")
	}
	return config, nil
}
