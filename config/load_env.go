package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"SQL_HOST"`
	DBUser     string `mapstructure:"SQL_USER"`
	DBPassword string `mapstructure:"SQL_PASSWORD"`
	DBName     string `mapstructure:"SQL_DB"`
	DBPort     string `mapstructure:"SQL_PORT"`

	RedisUrl string `mapstructure:"REDIS_URL"`
}

// LoadConfig loads the configuration from the specified path.
func LoadConfig(path string) (config Config, err error) {
	// Add the specified path to the list of configuration paths.
	viper.AddConfigPath(path)

	// Set the configuration type to "env".
	viper.SetConfigType("env")

	// Set the configuration name to "app".
	viper.SetConfigName("app")

	// Enable automatic environment variable substitution.
	viper.AutomaticEnv()

	// Read the configuration from the specified path.
	err = viper.ReadInConfig()

	// If there was an error reading the configuration, return the error.
	if err != nil {
		return
	}

	// Unmarshal the configuration into the specified config struct.
	err = viper.Unmarshal(&config)

	// Return the loaded configuration and any error that occurred.
	return
}
