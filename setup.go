package main

import (
	"fmt"
	"log"
	"mmacli/config"
)

var sampleNames = map[int]string{1: "Anderson Silva", 2: "Chael Sonnen"}
var Players map[int]*Fighter

func CreateFighter(index int) Fighter {
	fmt.Printf("Type the name of fighter %d: ", index)
	name := ""
	if *HideDelayFlag {
		name = sampleNames[index]
	} else {
		name = ReadString()
	}
	health := MaxHealth
	return Fighter{name, health}
}

func SetupPlayers() {
	fighterOne := CreateFighter(1)
	fighterTwo := CreateFighter(2)

	WriteString(fmt.Sprintf("\n== %s versus %s ==\n", fighterOne.name, fighterTwo.name))

	Players = map[int]*Fighter{1: &fighterOne, 2: &fighterTwo}
}

func SetupAttackList() {
	err := config.LoadAttackList("config/attacks.json")
	if err != nil {
		log.Fatal("Error on loading attack list:", err)
	}
}

func SetupGame() error {
	fmt.Printf("# Welcome to MMA CLI\n\n")

	SetupPlayers()
	SetupAttackList()

	return nil
}
