package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"mmacli/config/attack"
	"mmacli/config/defense"
	"mmacli/config/flags"
	"mmacli/myfmt"
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

	// Init delay callback to sync delay behavior
	myfmt.SetDelayCallback(TimeDelay)
}

func loadAssets() error {
	if err := flags.LoadList("config/flags.json"); err != nil {
		return fmt.Errorf("Error loading feature flags: %e", err)
	}
	if err := attack.LoadList("config/attacks.json"); err != nil {
		return fmt.Errorf("error on load attack list: %e", err)
	}
	if err := defense.LoadList("config/defenses.json"); err != nil {
		return fmt.Errorf("error on load defense list: %e", err)
	}

	return nil
}

func gameLoop() {
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

func main() {
	// Load assets
	if err := loadAssets(); err != nil {
		log.Fatal(err)
	}

	// Setup command-line flags
	flag.Parse()
	if *HideDelayFlag {
		Delay = 0
	}

	if err := SetupGame(); err != nil {
		log.Fatal("Error on setup game:", err)
	}

	gameLoop()
}
