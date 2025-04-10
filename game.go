package main

import (
	"fmt"
	"math/rand"
	"time"
)

var clock time.Duration = (60 + 60 + 30) * time.Second
var round int = 1

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
	// clock = clock.Add(time.Duration(roll) * time.Second)
	clock -= time.Duration(roll) * time.Second
}

func GetClockTime() string {
	clockInSecs := int(clock / time.Second)
	m := clockInSecs / 60
	s := clockInSecs - (m * 60)
	return fmt.Sprintf("%02d:%02d\n\n", m, s)
}

func GetPlayer(currentPlayerIndex int, players map[int]*Fighter) *Fighter {
	return players[currentPlayerIndex]
}

func GameLoop() {
	var currentPlayerIndex int
	shouldRunInitiative := true

	for CheckGameOver(Players) == false {
		PrintStatus(Players)
		if shouldRunInitiative {
			currentPlayerIndex = PlayersInitiative(Players)
			PressEnter()
			PrintStatus(Players)
		}

		opponentIndex := Swap[currentPlayerIndex]
		PlayerAttack(currentPlayerIndex, Players)
		TimeDelay()
		PressEnter()

		for range 5 {
			AddClockTime()
		}

		currentPlayerIndex = opponentIndex
		shouldRunInitiative = !shouldRunInitiative
	}

	fmt.Println("Game over")
}

func GetRound() string {
	return fmt.Sprintf("%d", round)
}

func PressEnter() {
	fmt.Printf("\n\nPress [ENTER] to continue.")
	if *NoPressEnterFlag {
		return
	}
	ReadChar()
}
