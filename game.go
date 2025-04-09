package main

import (
	"math/rand"
)

func CheckGameOver(players map[int]*Fighter) bool {
	for _, p := range players {
		if p.health <= 0 {
			return true
		}
	}
	return false
}

func RollDice() int {
	return rand.Intn(6) + 1
}

func GetPlayer(currentPlayerIndex int, players map[int]*Fighter) *Fighter {
	return players[currentPlayerIndex]
}
