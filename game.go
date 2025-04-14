package main

import (
	"fmt"
	"math/rand"
	"mmacli/models"
	"time"
)

var clock time.Duration = (60 + 60 + 30) * time.Second
var round int = 1

func CheckGameOver() bool {
	for _, p := range models.GetPlayers() {
		if p.Health <= 0 {
			return true
		}
	}
	return false
}

func RollDice() int {
	return rand.Intn(6) + 1
}

func AddClockTime() {
	for range 5 {
		roll := RollDice()
		clock -= time.Duration(roll) * time.Second
		if clock < 0 {
			clock = 0
		}
	}
}

func GetClockTime() string {
	clockInSecs := int(clock / time.Second)
	m := clockInSecs / 60
	s := clockInSecs - (m * 60)
	return fmt.Sprintf("%02d:%02d\n\n", m, s)
}

func GameLoop() {
	var (
		currentPlayerIndex  int
		isDraw              bool
		shouldRunInitiative = true
	)

	for CheckGameOver() == false {
		PrintStatus()
		AddClockTime()

		if shouldRunInitiative {
			currentPlayerIndex, isDraw = PlayersInitiative()
			PressEnter()
			ClearScreen()
			if isDraw {
				continue
			}
		}

		opponentIndex := Swap[currentPlayerIndex]
		PlayerAttack(currentPlayerIndex)
		currentPlayerIndex = opponentIndex
		shouldRunInitiative = !shouldRunInitiative

		TimeDelay()
		PressEnter()
	}

	fmt.Println("Game over")
}

func GetRound() string {
	return fmt.Sprintf("%d", round)
}

func PressEnter() {
	fmt.Printf("\n\nPress [ENTER] to continue.")
	if GetCliFlag("noenter") || GetCliFlag("n") {
		return
	}
	ReadChar()
}
