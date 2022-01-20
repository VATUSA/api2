package config

import (
	"io/ioutil"
	"os"

	"sigs.k8s.io/yaml"
)

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

	return &c, nil
}
