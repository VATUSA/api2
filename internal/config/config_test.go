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
    automigrate: true

redis:
    password: secret
    database: 0
    sentinel: true
    master_name: mymaster
    sentinel_addrs:
    - sentinel-0.redis.svc:26379
    address: redis.redis.svc:6379

session:
    cookie:
        name: vatusa
        secret: password
    jwt:
        jwks_path: jwks.json`)

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

		Redis: ConfigRedis{
			Password:   "secret",
			DB:         0,
			Sentinel:   true,
			MasterName: "mymaster",
			SentinelAddrs: []string{
				"sentinel-0.redis.svc:26379",
			},
			Address: "redis.redis.svc:6379",
		},

		Session: ConfigSession{
			Cookie: ConfigSessionCookie{
				Name:   "vatusa",
				Secret: "password",
				Domain: ".vatusa.net", // Default
				Path:   "/",           // Default
				MaxAge: 604800,        // Default
			},
			JWT: ConfigSessionJWT{
				JWKSPath: "jwks.json",
			},
		},
	}

	config, err := Parse(yamlConfig)
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}
	assert.Equal(t, structConfig, *config)
}
