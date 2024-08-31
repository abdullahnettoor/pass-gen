package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration values loaded from environment variables
type Config struct {
	DbConnectionString string `mapstructure:"DB_CONNECTION_URI"` // PostgreSQL connection string
	JwtSecret         string `mapstructure:"JWT_SECRET"`        // Secret key for JWT token generation
	CipherSecret      string `mapstructure:"CIPHER_SECRET"`     // Secret key for AES encryption
	ConfigFilePath    string `mapstructure:"CONFIG_FILE_PATH"`  // Path to store config file
	ConfigPath        string `mapstructure:"CONFIG_PATH"`       // Base path for config directory
}

// InitConfig loads configuration from dev.env file and environment variables
// Returns a pointer to Config struct and any error encountered
func InitConfig() (*Config, error) {
	var config = Config{}

	// Configure viper to read from dev.env file
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Read configuration file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Map configuration values to struct
	err = viper.Unmarshal(&config)
	return &config, err
}
