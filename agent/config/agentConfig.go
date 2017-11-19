package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Path to yaml file, used for server configuration
const (
	CfgPath           = "../../agent/agentServer/agentConfig.yaml"
	DefaultPort       = "9001"
	DefaultHost       = "localhost"
	DefaultServerAddr = "localhost:9000"
	Heartbeat         = 3
)

// Config structs yaml configuration
type Config struct {
	Port          string `yaml:"port"`
	Host          string `yaml:"localhost"`
	DynamicServer string `yaml:"mainServerAddr"`
	Heartbeat     int    `yaml:"heartbeat"`
}

var serverConfig *Config

func init() {
	serverConfig = GetSettings(CfgPath)
	serverConfig.setPort()
	serverConfig.setDynamicServer()
	serverConfig.setHeartbeat()
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
	if cfg.Port == "" {
		cfg.Port = DefaultPort
	}
}

func (cfg *Config) setDynamicServer() {
	if cfg.DynamicServer == "" {
		cfg.DynamicServer = DefaultServerAddr
	}
}

func (cfg *Config) setHeartbeat() {
	if cfg.Heartbeat == 0 {
		cfg.Heartbeat = Heartbeat
	}
}

// GetAddr returns address to listen on
func GetAddr() string {
	return serverConfig.Host + ":" + serverConfig.Port
}

// GetServerAddr returns address of main server
func GetServerAddr() string {
	return serverConfig.DynamicServer
}

// GetHeartbeat returns frequency of sending data to main server (in seconds)
func GetHeartbeat() int {
	return serverConfig.Heartbeat
}
