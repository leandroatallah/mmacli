package main

import (
	"fmt"
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

func GameLoop() {
	var currentPlayerIndex int
	shouldRunInitiative := true

	for CheckGameOver(Players) == false {
		TimeDelay()
		if shouldRunInitiative {
			currentPlayerIndex = PlayersInitiative(Players)
		}

		opponentIndex := Swap[currentPlayerIndex]
		PlayerAttack(currentPlayerIndex, Players)
		TimeDelay()

		PrintStatus(Players)

		currentPlayerIndex = opponentIndex
		shouldRunInitiative = !shouldRunInitiative
	}

	fmt.Println("Game over")
}
