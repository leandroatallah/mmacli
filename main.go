package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"mmacli/config"
	"mmacli/config/attack"
	"mmacli/config/defense"
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
	// Load assets
	err := config.LoadFeatureFlags("config/flags.json")
	if err != nil {
		log.Fatal("Error loading feature flags:", err)
	}
	err = attack.LoadList("config/attacks.json")
	if err != nil {
		log.Fatal("error on load attack list:", err)
	}
	err = defense.LoadList("config/defenses.json")
	if err != nil {
		log.Fatal("error on load defense list:", err)
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
