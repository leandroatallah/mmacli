package main

import (
	"flag"
	"log"
	"math/rand"
	"mmacli/myfmt"
	"time"
)

var Players map[int]*Fighter
var Delay = 1 * time.Second
var HideDelayFlag = flag.Bool("quick", false, "Skip delays")
var NoPressEnterFlag = flag.Bool("noenter", false, "Skip press enter command")
var UseMockFlag = flag.Bool("mock", false, "Use mock to fill players name")

type Fighter struct {
	name   string
	health int
}

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
