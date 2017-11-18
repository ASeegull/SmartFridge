package staticServerConfig

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Path to yaml file, used for server configuration
const (
	CfgPath           = "../../staticServer/staticServerConfig/clientConfig.yaml"
	DefaultPort       = "5080"
	DefaultStaticPath = "../../staticServer/static"
	DefaultPrefix     = "/static/"
	DefaultServerAddr = "http://localhost:9000"
)

// Config structs yaml configuration
type Config struct {
	Port            string `yaml:"port"`
	Host            string `yaml:"localhost"`
	StaticFilesPath string `yaml:"staticFilesPath"`
	Prefix          string `yaml:"pathPrefix"`
	DynamicServer   string `yaml:"dynamicServerAddress"`
}

var serverConfig *Config

func init() {
	serverConfig = GetSettings(CfgPath)
	serverConfig.setPath()
	serverConfig.setPort()
	serverConfig.setDynamicServer()
}

// GetSettings reads configuration file and stores values to struct variable
func GetSettings(cfgPath string) *Config {
	c := new(Config)
	cfg, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Errorf("Failed to read config file %v", err)
		return c
	}
	err = yaml.Unmarshal(cfg, c)
	if err != nil {
		log.Errorf("Failed to read config file %v", err)
	}
	return c
}

// GetAddr sets port for dev or production
func (cfg *Config) setPort() {
	// Checks if system sets port (like on Heroku)
	port := os.Getenv("PORT")
	if port != "" {
		cfg.Port = port
		return
	}
	// Sets port to default when none specified
	if port == "" && cfg.Port == "" {
		cfg.Port = DefaultPort
	}
}

// setPath sets path to static files to default values if none specified in configuration file
func (cfg *Config) setPath() {
	if cfg.StaticFilesPath == "" {
		cfg.StaticFilesPath = DefaultStaticPath
	}

	if cfg.Prefix == "" {
		cfg.Prefix = DefaultPrefix
	}
}

func (cfg *Config) setDynamicServer() {
	if cfg.DynamicServer == "" {
		cfg.DynamicServer = DefaultServerAddr
	}
}

// GetAddr returns address to listen on
func GetAddr() string {
	return serverConfig.Host + ":" + serverConfig.Port
}

// GetStaticPath returns location of static Files
func GetStaticPath() (path, prefix string) {
	return serverConfig.StaticFilesPath, serverConfig.Prefix
}

// GetServerAddr returns address of main server
func GetServerAddr() string {
	return serverConfig.DynamicServer
}
