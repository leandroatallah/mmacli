package main

import (
	"fmt"
	"mmacli/fight"
	"mmacli/models"
	"mmacli/utils"
	"time"
)

var clock time.Duration = (60 + 60 + 30) * time.Second
var round int = 1

func checkGameOver() bool {
	for _, p := range models.GetPlayers() {
		if p.Health <= 0 {
			return true
		}
	}
	return false
}

func addClockTime() {
	for range 5 {
		roll := utils.RollDice()
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

	for checkGameOver() == false {
		PrintStatus()
		addClockTime()

		if shouldRunInitiative {
			currentPlayerIndex, isDraw = PlayersInitiative()
			PressEnter()
			if isDraw {
				continue
			}
		}

		PrintStatus()
		opponentIndex := utils.Swap[currentPlayerIndex]
		fight.PlayerAttack(currentPlayerIndex)
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
