package config

import (
	"encoding/json"
	"os"
)

type FeatureFlags map[string]bool

func LoadFeatureFlags(filename string) (FeatureFlags, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var flags FeatureFlags
	err = json.Unmarshal(file, &flags)
	if err != nil {
		return nil, err
	}

	return flags, nil
}
