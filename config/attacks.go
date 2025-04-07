package config

import (
	"encoding/json"
	"os"
	"sort"
)

type Attack struct {
	Index       int    `json:"index"`
	Name        string `json:"name"`
	AttackBonus int    `json:"attackBonus"`
}

var attackList []Attack

func LoadAttackList(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &attackList)
	if err != nil {
		return err
	}

	sort.Slice(attackList, func(i, j int) bool {
		return attackList[i].Index < attackList[j].Index
	})

	return nil
}

func FindAttackByIndex(index int) int {
	for i, attack := range attackList {
		if attack.Index == index {
			return i
		}
	}

	return -1
}

func GetAttackByIndex(index int) (*Attack, bool) {
	attackIndex := FindAttackByIndex(index)
	if attackIndex == -1 {
		return nil, false
	}
	attack := attackList[attackIndex]

	return &attack, true
}

func GetAllAttackList() []Attack {
	return attackList
}
