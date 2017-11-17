package staticServer

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Path to yaml file, used for server configuration
const (
	CfgPath           = "../../clientConfig.yaml"
	DefaultPort       = ":5080"
	DefaultStaticPath = "../../staticServer/static"
	DefaultPrefix     = "/static/"
)

// Saves configs, extracted from .yaml file
var serverCfg *config


type config struct {
	Port         string `yaml:"port"`
	ServerAddres string `yaml:"serverAddress"`
	Host         string `yaml:"localhost"`
	StaticFilesPath string `yaml:"staticFilesPath"`
	Prefix          string `yaml:"pathPrefix"`
}

// Reads configuration file and stores values to struct variable
func getSettings(cfgPath string) *config {
	cfg, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.WithField("error", err).Error("Failed to read config file")
	}
	c := new(config)
	rerr := yaml.Unmarshal(cfg, c)
	if err != nil {
		log.WithField("error", rerr).Error("Failed to read config file")
	}
	return c
}

// GetPort sets port for dev or production
func GetPort() string {
	// Checks if system sets port (like on Heroku)
	port := ":" + os.Getenv("PORT")
	// Then looks if it's set in config file 
	if port == ":" {
		port = serverCfg.Port
	}
	// If none of above works, sets port to default
	if port == ":" {
		port = DefaultPort
	}
	return port
}

// Sets path to static files to default values if none specified in configuration file
func getPath(cfg *config) (path, prefix string) {
	if cfg.StaticFilesPath == "" {
		path = DefaultStaticPath
	}
	if cfg.Prefix == "" {
		prefix = DefaultPrefix
	}
	return
}

// NewRouter сreates and returns gorilla router
func NewRouter() *mux.Router {
	filePath, prefix := getPath(serverCfg)
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(filePath)))
	r.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(http.Dir(filePath))))
	r.Path("/content").HandlerFunc(fetch("/client/fridgeContent")).Methods("GET")
	r.Path("/recipes").HandlerFunc(fetch("/client/allRecipes")).Methods("GET")
	return r
}

func fetch(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(serverCfg.ServerAddres + query)
		if err != nil {
			log.Error("Could not get response", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
		defer res.Body.Close()
		
		jsn, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.WithField("error", err).Error("Could not get response because of")
		}

		done, err := w.Write(jsn)
		if err != nil {
			log.WithField("error", err).Error("Could not write response to client because of")
		}
		log.Infof("Success, written down %d bytes", done)
	}
}


func init() {
	serverCfg = getSettings(CfgPath)
}
