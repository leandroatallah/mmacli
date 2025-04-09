package defense

import (
	"encoding/json"
	"os"
	"sort"
)

type Defense struct {
	Index    int    `json:"index"`
	Name     string `json:"name"`
	Accuracy int    `json:"accuracy"`
}

var defenseList []Defense

func LoadList(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &defenseList)
	if err != nil {
		return err
	}

	sort.Slice(defenseList, func(i, j int) bool {
		return defenseList[i].Index < defenseList[j].Index
	})

	return nil
}

func FindByIndex(index int) int {
	for i, attack := range defenseList {
		if attack.Index == index {
			return i
		}
	}

	return -1
}

func GetByIndex(index int) (*Defense, bool) {
	attackIndex := FindByIndex(index)
	if attackIndex == -1 {
		return nil, false
	}
	attack := defenseList[attackIndex]

	return &attack, true
}

func GetAll() []Defense {
	return defenseList
}
