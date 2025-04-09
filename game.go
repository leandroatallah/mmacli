package main

import (
	"fmt"
	"math/rand"
	"mmacli/myfmt"
	"time"
)

var Clock time.Time

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

func AddClockTime() {
	roll := RollDice()
	// TODO: Limit to round time limit
	Clock = Clock.Add(time.Duration(roll) * time.Second)
}

func GetPlayer(currentPlayerIndex int, players map[int]*Fighter) *Fighter {
	return players[currentPlayerIndex]
}

func GameLoop() {
	var currentPlayerIndex int
	shouldRunInitiative := true

	for CheckGameOver(Players) == false {
		formattedClock := Clock.Format("15:04:05")
		myfmt.PrintDelay("%v\n\n", formattedClock[3:])
		if shouldRunInitiative {
			currentPlayerIndex = PlayersInitiative(Players)
		}

		opponentIndex := Swap[currentPlayerIndex]
		PlayerAttack(currentPlayerIndex, Players)
		TimeDelay()

		for range 5 {
			AddClockTime()
		}

		PrintStatus(Players)

		currentPlayerIndex = opponentIndex
		shouldRunInitiative = !shouldRunInitiative

	}

	fmt.Println("Game over")
}
