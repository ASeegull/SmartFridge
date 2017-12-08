package config

import (
	"io/ioutil"

	"github.com/davecheney/errors"
	yaml "gopkg.in/yaml.v2"
)

// config is a struct yaml configuration
type Config struct {
	Host          string `yaml:"host"`
	HTTPSPort     string `yaml:"HTTPSport"`
	HTTPPort      string `yaml:"HTTPport"`
	StaticPath    string `yaml:"staticPatn"`
	ServerAddress string `yaml:"serverAddress"`
	Cert          string `yaml:"cert"`
	Key           string `yaml:"key"`
}

// GetSettings reads configuration file and stores values to struct variable
func GetSettings(cfgPath string) (*Config, error) {
	read, err := ioutil.ReadFile(cfgPath)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to read config file")
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(read, &cfg); err != nil {
		return nil, errors.Annotate(err, "Failed to unmarshal config file")
	}
	return cfg, err
}

// HTTPAddr returns address for HTTP server to listen on
func (cfg *Config) HTTPAddr() string {
	return cfg.Host + ":" + cfg.HTTPPort
}

// HTTPSAddr returns address for HTTPS server to listen on
func (cfg *Config) HTTPSAddr() string {
	return cfg.Host + ":" + cfg.HTTPSPort
}
