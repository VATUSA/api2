package seed

import (
	"encoding/json"
	"errors"

	"sigs.k8s.io/yaml"
)

type Seed struct {
	Kind   string                   `json:"kind"`
	Values []map[string]interface{} `json:"values"`
}

func IsValidSeedType(seedType string) bool {
	return seedType == "yaml" || seedType == "json"
}

func IsValidSeedKind(table string) bool {
	return table == "user" || table == "rating" || table == "facility"
}

func BuildSeed(t string, data []byte) (*Seed, error) {
	var seed Seed

	if !IsValidSeedType(t) {
		return nil, errors.New("invalid seed type " + t)
	}

	switch t {
	case "yaml":
		err := yaml.Unmarshal(data, &seed)
		if err != nil {
			return nil, err
		}
	case "json":
		err := json.Unmarshal(data, &seed)
		if err != nil {
			return nil, err
		}
	}

	if !IsValidSeedKind(seed.Kind) {
		return nil, errors.New("invalid seed kind " + seed.Kind)
	}

	return &seed, nil
}
