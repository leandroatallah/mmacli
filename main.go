package main

import (
	"log"
	"math/rand"
	"mmacli/myfmt"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// Init delay callback to sync delay behavior
	myfmt.SetDelayCallback(TimeDelay)
}

func main() {
	ClearScreen()
	if err := SetupGame(); err != nil {
		log.Fatal("Error on setup game:", err)
	}

	GameLoop()
}
