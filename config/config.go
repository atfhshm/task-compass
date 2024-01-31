package config

import (
	"errors"
	"os"
)

var (
	ErrDbUrl      = errors.New("unable to find DB_URL in environment variables")
	ErrServerPort = errors.New("unable to find PORT in environment variables")
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

func NewServerConfig() (*ServerConfig, error) {
	var serverConfig ServerConfig
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return nil, ErrServerPort
	}
	serverConfig.Port = port
	return &serverConfig, nil
}

type DatabaseConfig struct {
	URL string
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	var databaseConfig DatabaseConfig
	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		return nil, ErrDbUrl
	}
	databaseConfig.URL = dbUrl
	return &databaseConfig, nil
}

func NewConfig() (*Config, error) {
	var config Config

	serverConfig, err := NewServerConfig()
	if err != nil {
		return nil, err
	}

	databaseConfig, err := NewDatabaseConfig()
	if err != nil {
		return nil, err
	}

	config.Server = *serverConfig
	config.Database = *databaseConfig
	return &config, nil
}
