package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type MongoConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
	Table    string `yaml:"table"`
}

type ServerConfig struct {
	Port            string `yaml:"port"`
	Host            string `yaml:"host"`
	ReadBufferSize  int    `yaml:"readBufferSize"`
	WriteBufferSize int    `yaml:"writeBufferSize"`
}

const (
	mongoConfigPath  = "../SmartFridge/server/config/mongoConfig.yaml"
	serverConfigPath = "../SmartFridge/server/config/fridgeServerConfig.yaml"
)

func GetMongoConfig() *MongoConfig {
	config := &MongoConfig{}
	yamlFile, err := ioutil.ReadFile(mongoConfigPath)
	if err != nil {
		log.Error(err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Error(err)
	}

	return config
}

func GetServerConfig() *ServerConfig {
	config := &ServerConfig{}
	yamlFile, err := ioutil.ReadFile(serverConfigPath)
	if err != nil {
		log.Error(err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Error(err)
	}

	return config
}
