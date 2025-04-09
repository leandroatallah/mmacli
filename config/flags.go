package config

import (
	"encoding/json"
	"os"
)

type FeatureFlags map[string]bool

// TODO: Change to a local variable and use a getter
var Flags FeatureFlags

func LoadFeatureFlags(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Flags)
	if err != nil {
		return err
	}

	return nil
}
