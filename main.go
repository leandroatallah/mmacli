package main

import (
	"flag"
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

var Delay = 1 * time.Second

func main() {
	hideDelay := flag.Bool("quick", false, "Skip delays")
	flag.Parse()

	if *hideDelay {
		Delay = 0
	}

	fmt.Printf("# Welcome to MMA CLI\n\n")

	fighterOne := SetupPlayer(1)
	fighterTwo := SetupPlayer(2)

	WriteString(fmt.Sprintf("\n== %s versus %s ==\n", fighterOne.name, fighterTwo.name))

	players := map[int]*Fighter{1: &fighterOne, 2: &fighterTwo}
	var currentPlayerIndex int
	shouldRunInitiative := true

	for CheckGameOver(players) == false {
		time.Sleep(Delay)
		if shouldRunInitiative {
			currentPlayerIndex = PlayersInitiative(players)
		}

		opponentIndex := Swap[currentPlayerIndex]
		PlayerAttack(currentPlayerIndex, players)
		time.Sleep(Delay)
		fmt.Println()

		PrintStatus(&fighterOne, &fighterTwo)

		currentPlayerIndex = opponentIndex

		shouldRunInitiative = !shouldRunInitiative
	}

	fmt.Println("Game over")
}
