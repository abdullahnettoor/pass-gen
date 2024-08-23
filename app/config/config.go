package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbConnectionString string `mapstructure:"DB_CONNECTION_URI"`
	JwtSecret          string `mapstructure:"JWT_SECRET"`
	ConfigFilePath     string `mapstructure:"CONFIG_FILE_PATH"`
	ConfigPath         string `mapstructure:"CONFIG_PATH"`
}

func InitConfig() (*Config, error) {
	var config = Config{}

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return &config, err
}
