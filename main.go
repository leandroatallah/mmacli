package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"mmacli/config"
	"time"
)

type Fighter struct {
	name   string
	health int
}

var Delay = 1 * time.Second
var HideDelayFlag = flag.Bool("quick", false, "Skip delays")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Load feature flags
	err := config.LoadFeatureFlags("config/flags.json")
	if err != nil {
		log.Fatal("Error loading feature flags:", err)
	}

	// Setup command-line flags
	flag.Parse()
	if *HideDelayFlag {
		Delay = 0
	}

	err = SetupGame()
	if err != nil {
		log.Fatal("Error on setup game:", err)
	}

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
