package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

func NewServerConfig() (*ServerConfig, error) {
	var serverConfig ServerConfig
	port, ok := os.LookupEnv("PORT")
	if !ok {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
		port = os.Getenv("PORT")
	}
	serverConfig.Port = port
	return &serverConfig, nil
}

type DatabaseConfig struct {
	URL string `json:"url"`
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	var databaseConfig DatabaseConfig
	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
		dbUrl = os.Getenv("PORT")
	}
	databaseConfig.URL = dbUrl
	return &databaseConfig, nil
}

func NewConfig() (*Config, error) {
	var config Config

	var serverConfig *ServerConfig
	var serverErr error

	var databaseConfig *DatabaseConfig
	var databaseErr error

	serverConfig, serverErr = NewServerConfig()
	if serverErr != nil {
		return nil, serverErr
	}

	databaseConfig, databaseErr = NewDatabaseConfig()
	if databaseErr != nil {
		return nil, databaseErr
	}

	config.Server = *serverConfig
	config.Database = *databaseConfig
	return &config, nil
}
