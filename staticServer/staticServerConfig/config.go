package staticServerConfig

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/davecheney/errors"
	yaml "gopkg.in/yaml.v2"
)

// Path to yaml file, used for server configuration
const (
	CfgPath = "./clientConfig.yaml"
)

var config *Config

func init() {
	var err error
	config, err = getSettings(CfgPath)
	if err != nil {
		log.Fatal(err)
	}
}

// Config structs yaml configuration
type Config struct {
	HTTPcfg       `yaml:",inline"`
	HTTPScfg      `yaml:",inline"`
	Static        `yaml:",inline"`
	ServerAddress string `yaml:"serverAddress"`
	Host          string `yaml:"host"`
}

type HTTPcfg struct {
	HTTPPort string `yaml:"HTTPport"`
}

type HTTPScfg struct {
	HTTPSPort string `yaml:"HTTPSport"`
	Cert      string `yaml:"cert"`
	Key       string `yaml:"key"`
}

type Static struct {
	Path   string `yaml:"staticFilesPath"`
	Prefix string `yaml:"pathPrefix"`
}

// GetSettings reads configuration file and stores values to struct variable
func getSettings(cfgPath string) (*Config, error) {
	read, err := ioutil.ReadFile(cfgPath)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to read config file")
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(read, &cfg); err != nil {
		return nil, errors.Annotate(err, "Failed to read config file %v")
	}
	return cfg, err
}

func GetStaticFilesPath() *Static {
	return &config.Static
}

func GetStaticHTTPScfg() *HTTPScfg {
	return &config.HTTPScfg
}

func GetDynamicServer() string {
	return config.ServerAddress
}

func GetHTTPAddr() string {
	return config.Host + ":" + config.HTTPcfg.HTTPPort
}

func GetHTTPSAddr() string {
	return config.Host + ":" + config.HTTPScfg.HTTPSPort
}

func GetPem() (string, string) {
	return config.HTTPScfg.Cert, config.HTTPScfg.Key
}
