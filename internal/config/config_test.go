package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigParse(t *testing.T) {
	yamlConfig := []byte(`
server:
    port: 3000
    host: 0.0.0.0

database:
    host: 127.0.0.1
    port: 5432
    user: root
    password: secret12345
    database: vatusa_dev
    driver: pgsql
    automigrate: true`)

	structConfig := Config{
		Server: ConfigServer{
			Port: "3000",
			Host: "0.0.0.0",
		},
		Database: ConfigDatabase{
			Host:        "127.0.0.1",
			Port:        "5432",
			User:        "root",
			Password:    "secret12345",
			Database:    "vatusa_dev",
			Driver:      "pgsql",
			AutoMigrate: true,
		},
	}

	config, err := Parse(yamlConfig)
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}
	assert.Equal(t, structConfig, *config)
}
