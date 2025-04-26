package configs

import (
	"errors"
	"os"
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
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	PassSecret    string
}

func loadConfig() (*config, error) {
	cfg := &config{
		DBDriver:      os.Getenv("GymControll_DBDRIVER"),
		DBHost:        os.Getenv("GymControll_DBHOST"),
		DBPort:        os.Getenv("GymControll_DBPORT"),
		DBUser:        os.Getenv("GymControll_DBUSER"),
		DBPassword:    os.Getenv("GymControll_DBPASS"),
		DBName:        os.Getenv("GymControll_DBNAME"),
		WebServerPort: os.Getenv("GymControll_WEBSERVER_PORT"),
		PassSecret:    os.Getenv("GymControll_PassSecret"),
	}

	if cfg.DBDriver == "" || cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
		return nil, errors.New("missing required environment variables")
	}

	return cfg, nil
}
