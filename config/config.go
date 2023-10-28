package config

import (
	"github.com/spf13/viper"
)

type InternalConfig struct {
	Port int
}

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Config struct {
	InternalConfig *InternalConfig
	MySQLConfig    *MySQLConfig
}

func Get() *Config {
	viper.AutomaticEnv()

	return &Config{
		InternalConfig: &InternalConfig{
			Port: viper.GetInt("SERVER_PORT"),
		},
		MySQLConfig: &MySQLConfig{
			Host:     viper.GetString("MYSQL_HOST"),
			Port:     viper.GetInt("MYSQL_PORT"),
			Username: viper.GetString("MYSQL_USERNAME"),
			Password: viper.GetString("MYSQL_PASSWORD"),
			Database: viper.GetString("MYSQL_DATABASE"),
		},
	}
}
