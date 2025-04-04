package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Fighter struct {
	name   string
	health int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Welcome to MMA CLI")

	fighterOne := SetupPlayer(1)
	fighterTwo := SetupPlayer(2)

	WriteString(fmt.Sprintf("%s versus %s\n", fighterOne.name, fighterTwo.name))

	players := map[int]*Fighter{1: &fighterOne, 2: &fighterTwo}
	var currentPlayerIndex int
	shouldRunInitiative := true

	for CheckGameOver(players) == false {
		time.Sleep(time.Second * 1)
		if shouldRunInitiative {
			currentPlayerIndex = PlayersInitiative(players)
		}

		opponentIndex := Swap[currentPlayerIndex]
		PlayerAttack(currentPlayerIndex, players)
		time.Sleep(time.Second * 1)
		fmt.Println()

		PrintStatus(&fighterOne, &fighterTwo)

		currentPlayerIndex = opponentIndex

		shouldRunInitiative = !shouldRunInitiative
	}

	fmt.Println("Game over")
}
