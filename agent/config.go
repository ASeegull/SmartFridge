package agent

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
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
		return nil, errors.Wrapf(err, "could not read yaml file %s", configPath)
	}

	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return nil, errors.Wrapf(err, "could not decode config file %s", configPath)
	}
	return config, nil
}

// GetEndPoints returns API endpoint to call for setup and address to call to establish websocket connection
func (cfg *Config) GetEndPoints() string {
	return fmt.Sprintf("%s:%s/%s", cfg.Websocket, cfg.Port, cfg.RestURI)
}
