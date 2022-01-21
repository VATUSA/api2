package config

import (
	"io/ioutil"
	"os"

	"github.com/imdario/mergo"
	"sigs.k8s.io/yaml"
)

var Cfg *Config

var defaultConfig = Config{
	Server: ConfigServer{
		Port: "3000",
		Host: "0.0.0.0",
	},
	Database: ConfigDatabase{
		Host:        "127.0.0.1",
		Port:        "5432",
		User:        "postgres",
		Password:    "secret12345",
		Database:    "vatusa",
		Driver:      "postgres",
		AutoMigrate: true,
	},
	Redis: ConfigRedis{
		Password: "",
		DB:       0,
		Address:  "127.0.0.1:6379",
	},
	Session: ConfigSession{
		Cookie: ConfigSessionCookie{
			Name:   "vatusa",
			Secret: "secret12345",
			Domain: ".vatusa.net",
			Path:   "/",
			MaxAge: 604800,
		},
		JWT: ConfigSessionJWT{
			JWKSPath: "jwks.json",
		},
	},
}

func Load(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return Parse(data)
}

func Parse(config []byte) (*Config, error) {
	var c Config

	err := yaml.Unmarshal(config, &c)
	if err != nil {
		return nil, err
	}

	if err := mergo.Merge(&c, defaultConfig); err != nil {
		return nil, err
	}

	return &c, nil
}
