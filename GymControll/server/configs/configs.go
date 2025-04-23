package configs

import (
	"errors"
	"os"
	"strings"
)

var (
	Config *config
	err    error
)

func init() {
	Config, err = loadConfig()
	if err != nil {
		panic(errors.New("Error loading environment variables: " + err.Error()))
	}
}

type config struct {
	DBDriver         string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	WebServerPort    string
	GRPCServerPort   string
	MQServerHost     string
	MQUser           string
	MQPassword       string
	SPEXApplications []string
}

func loadConfig() (*config, error) {
	cfg := &config{
		DBDriver:       os.Getenv("SPEX_DBDRIVER"),
		DBHost:         os.Getenv("SPEX_DBHOST"),
		DBPort:         os.Getenv("SPEX_DBPORT"),
		DBUser:         os.Getenv("SPEX_DBUSER"),
		DBPassword:     os.Getenv("SPEX_DBPASS"),
		DBName:         os.Getenv("SPEX_DBNAME"),
		WebServerPort:  os.Getenv("SPEX_WEBSERVER_PORT"),
		GRPCServerPort: os.Getenv("SPEX_GRPCSERVER_PORT"),
		MQServerHost:   os.Getenv("SPEX_MQSERVER_HOST"),
		MQUser:         os.Getenv("SPEX_MQUSER"),
		MQPassword:     os.Getenv("SPEX_MQPASSWORD"),
	}

	applications := os.Getenv("SPEX_APPLICATIONS")
	if applications != "" {
		cfg.SPEXApplications = strings.Split(applications, ",")
	}

	if cfg.DBDriver == "" || cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
		return nil, errors.New("missing required environment variables")
	}

	return cfg, nil
}