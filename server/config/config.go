package config

import (
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//Config includes server and mongoDB configs
type Config struct {
	Server ServerConfig `yaml:"serverConfig"`
	Mongo  MongoConfig  `yaml:"mongoConfig"`
}

//MongoConfig includes config for mongoDB
type MongoConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
	Table    string `yaml:"table"`
}

//ServerConfig includes config for server
type ServerConfig struct {
	Port            string        `yaml:"port"`
	Host            string        `yaml:"host"`
	ReadBufferSize  int           `yaml:"readBufferSize"`
	WriteBufferSize int           `yaml:"writeBufferSize"`
	WebsocketSleep  time.Duration `yaml:"websocketSleep"`
}

const (
	serverConfigPath = "../SmartFridge/server/config/config.yaml"
)

var config *Config

//ReadConfig reads config from file
func ReadConfig() error {
	config = &Config{}
	yamlFile, err := ioutil.ReadFile(serverConfigPath)
	if err != nil {
		log.Error(err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

//GetServerConfig returns server config
func GetServerConfig() *ServerConfig {
	return &config.Server
}

//GetMongoConfig returns mongoDB config
func GetMongoConfig() *MongoConfig {
	return &config.Mongo
}
