package flags

import (
	"encoding/json"
	"os"
)

type featureFlags map[string]bool

// TODO: Change to a local variable and use a getter
var flags featureFlags

func LoadList(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &flags)
	if err != nil {
		return err
	}

	return nil
}

func GetById(flagId string) bool {
	flag, exists := flags[flagId]
	if !exists {
		return false
	}

	return flag
}
