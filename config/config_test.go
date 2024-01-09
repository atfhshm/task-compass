package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatabaseConfig(t *testing.T) {
	dbConfig, err := NewDatabaseConfig()
	assert.Nil(t, err, err)
	assert.NotNil(t, dbConfig)
}

func TestNewServerConfig(t *testing.T) {
	serverConfig, err := NewServerConfig()
	assert.Nil(t, err, err)
	assert.NotNil(t, serverConfig)
}

func TestNewConfig(t *testing.T) {
	conf, err := NewConfig()
	assert.Nil(t, err, err.Error())
	assert.NotNil(t, conf)
}
