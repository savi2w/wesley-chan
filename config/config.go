package config

import (
	"github.com/spf13/viper"
)

type AWSConfig struct {
	FileBucketName string
	Region         string
}

type InternalConfig struct {
	Port        int
	ServiceName string
}

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Config struct {
	AWSConfig      *AWSConfig
	InternalConfig *InternalConfig
	MySQLConfig    *MySQLConfig
}

func Get() *Config {
	viper.AutomaticEnv()

	return &Config{
		AWSConfig: &AWSConfig{
			FileBucketName: viper.GetString("AWS_FILE_BUCKET_NAME"),
			Region:         viper.GetString("AWS_REGION"),
		},
		InternalConfig: &InternalConfig{
			Port:        viper.GetInt("SERVER_PORT"),
			ServiceName: viper.GetString("SERVICE_NAME"),
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
