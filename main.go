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
var Flags config.FeatureFlags

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Load feature flags
	loadedflags, err := config.LoadFeatureFlags("config/flags.json")
	if err != nil {
		log.Fatal("Error loading feature flags:", err)
	}
	Flags = loadedflags

	// Setup command-line flags
	flag.Parse()
	if *HideDelayFlag {
		Delay = 0
	}

	players := SetupGame()
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

		PrintStatus(players)

		currentPlayerIndex = opponentIndex
		shouldRunInitiative = !shouldRunInitiative
	}

	fmt.Println("Game over")
}
