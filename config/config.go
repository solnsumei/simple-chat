package config

import "github.com/spf13/viper"

// Config struct
type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	SECRET     string
	PORT       string
}

// LoadConfigVars -- load confiuration
func LoadConfigVars() (*Config, error) {
	// Set environment file
	viper.SetConfigFile(".env")

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		return new(Config), err
	}

	config := &Config{
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		SECRET:     viper.GetString("SECRET"),
		PORT:       viper.GetString("PORT"),
	}

	return config, nil
}
