package configs

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Name        string
	Version     string
	Description string
	Host        string
	Port        int
}

type DatabaseConfig struct {
	Item      string
	Password  string
	Host      string
	Name      string
	Port      string
	Migration string
}

func NewConfig() *Config {
	config, err := ReadInConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func ReadInConfig() (config *Config, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("env.yml")
	viper.AddConfigPath("./configs")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
