package configs

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Config *config

func init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	Config, err = loadConfig(filepath.Join(configDir, "GymControl"))
	if err != nil {
		panic(errors.New("Error reading config file, " + err.Error()))
	}
}

type config struct {
	WebServerPort string `mapstructure:"GYMCONTROL_WEBSERVER_PORT"`
	DBServer      string `mapstructure:"GYMCONTROL_DB_SERVER"`
	DBServerPort  string `mapstructure:"GYMCONTROL_DB_SERVERPORT"`
	DBUser        string `mapstructure:"GYMCONTROL_DB_USER"`
	DBPassword    string `mapstructure:"GYMCONTROL_DB_PASSWORD"`
	DBDataBase    string `mapstructure:"GYMCONTROL_DB_DATABASE"`
}

func loadConfig(path string) (*config, error) {
	v := viper.New()
	var cfg *config
	v.SetConfigFile(".deploy/conf.toml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.New("Error reading config file, " + err.Error())
	}
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, errors.New("Error unmarshalling config file, " + err.Error())
	}
	return cfg, err
}
