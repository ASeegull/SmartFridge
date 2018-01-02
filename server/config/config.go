package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Config includes server and mongoDB configs
type Config struct {
	Server   ServerConfig      `yaml:"serverConfig"`
	Mongo    MongoConfig       `yaml:"mongoConfig"`
	Postgres PostgresConfigStr `yaml:"postgresConfig"`
}

//MongoConfig includes config for mongoDB
type MongoConfig struct {
	URI             string `yaml:"uri"`
	Database        string `yaml:"database"`
	Table           string `yaml:"table"`
	ConnectionsPool int    `yaml:"connectionsPool"`
}

//PostgresConfigStr structs yaml configuration
type PostgresConfigStr struct {
	Dbhost                   string `yaml:"dbhost"`
	Dbport                   string `yaml:"dbport"`
	DbUser                   string `yaml:"dbUser"`
	DbPassword               string `yaml:"dbPassword"`
	DbName                   string `yaml:"dbName"`
	MaxOpenedConnectionsToDb int    `yaml:"maxOpenedConnectionsToDb"`
	MaxIdleConnectionsToDb   int    `yaml:"maxIdleConnectionsToDb"`
	MbConnMaxLifetimeMinutes int    `yaml:"mbConnMaxLifetimeMinutes"`
}

//ServerConfig includes config for server
type ServerConfig struct {
	Port            string `yaml:"port"`
	Host            string `yaml:"host"`
	ReadBufferSize  int    `yaml:"readBufferSize"`
	WriteBufferSize int    `yaml:"writeBufferSize"`
}

const (
	serverConfigPath = "./config/config.yaml"
)

//ReadConfig reads config from file
func ReadConfig() (*Config, error) {
	config := &Config{}
	yamlFile, err := ioutil.ReadFile(serverConfigPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

////GetServerConfig returns server config
//func (c *Config) GetServerConfig() *ServerConfig {
//	return &c.Server
//}
//
////GetMongoConfig returns mongoDB config
//func (c *Config) GetMongoConfig() *MongoConfig {
//	return &c.Mongo
//}
