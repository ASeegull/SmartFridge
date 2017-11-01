package database

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	// Path to yaml file, used for postgres db configuration
	cfgPath = "./postgresConfig.yaml"
	//postgres connection credentials
	defaultdbhost     = "localhost"
	defaultdbport     = "5432"
	defaultdbUser     = "goproject"
	defaultdbPassword = ""
	defaultdbName     = "goproject"
)

//PostgresConfigStr structs yaml configuration
type PostgresConfigStr struct {
	Dbhost     string `yaml:"dbhost"`
	Dbport     string `yaml:"dbport"`
	DbUser     string `yaml:"dbUser"`
	DbPassword string `yaml:"dbPassword"`
	DbName     string `yaml:"dbName"`
}

var postgresConfig *PostgresConfigStr

func init() {
	var err error
	postgresConfig, err = GetPostgresConfig(cfgPath)
	if err != nil {
		log.Errorf("Failed to read config file %v", err)
	}
	postgresConfig.setDbhost()
	postgresConfig.setDbport()
	postgresConfig.setDbUser()
	postgresConfig.setDbPassword()
	postgresConfig.setDbName()

}

//GetPostgresConfig reads configuration file and stores values to struct variable
func GetPostgresConfig(cfgPath string) (*PostgresConfigStr, error) {
	c := new(PostgresConfigStr)
	cfg, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(cfg, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getDbhost() string {
	return postgresConfig.Dbhost
}

func getDbport() string {
	return postgresConfig.Dbport
}

func getDbUser() string {
	return postgresConfig.DbUser
}

func getDbPassword() string {
	return postgresConfig.DbPassword
}

func getDbName() string {
	return postgresConfig.DbName
}

// GetAddr sets port for dev or production
func (cfg *PostgresConfigStr) setDbhost() {
	if cfg.Dbhost == "" {
		cfg.Dbhost = defaultdbhost
	}
}

func (cfg *PostgresConfigStr) setDbport() {
	if cfg.Dbport == "" {
		cfg.Dbport = defaultdbport
	}
}

func (cfg *PostgresConfigStr) setDbUser() {
	if cfg.DbUser == "" {
		cfg.DbUser = defaultdbUser
	}
}

func (cfg *PostgresConfigStr) setDbName() {
	if cfg.DbName == "" {
		cfg.DbName = defaultdbName
	}
}

func (cfg *PostgresConfigStr) setDbPassword() {
	if cfg.DbPassword == "" {
		cfg.DbPassword = defaultdbPassword
	}
}
