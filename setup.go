package main

import "fmt"

var sampleNames = map[int]string{1: "Anderson Silva", 2: "Chael Sonnen"}

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

func SetupPlayers() map[int]*Fighter {
	fighterOne := CreateFighter(1)
	fighterTwo := CreateFighter(2)

	WriteString(fmt.Sprintf("\n== %s versus %s ==\n", fighterOne.name, fighterTwo.name))

	players := map[int]*Fighter{1: &fighterOne, 2: &fighterTwo}
	return players
}

func SetupGame() map[int]*Fighter {
	fmt.Printf("# Welcome to MMA CLI\n\n")

	players := SetupPlayers()

	return players
}
