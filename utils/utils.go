package utils

import (
	"math/rand"
)

var Swap = map[int]int{1: 2, 2: 1}

func RollDice() int {
	return rand.Intn(6) + 1
}
