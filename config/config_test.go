package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) SetupSuite() {
	log.Println("loading environment variable")
	godotenv.Load()
}

func (s *ConfigTestSuite) TestNewDatabaseConfig() {
	dbConfig, err := NewDatabaseConfig()
	s.Nil(err, err.Error())
	s.NotNil(dbConfig)
}

func (s *ConfigTestSuite) TestNewServerConfig() {
	serverConfig, err := NewServerConfig()
	s.Nil(err, err.Error())
	s.NotNil(serverConfig)
}

func (s *ConfigTestSuite) TestNewConfig() {
	conf, err := NewConfig()
	s.Nil(err, err.Error())
	s.NotNil(conf)
}
