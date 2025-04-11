package main

import (
	"flag"
	"fmt"
	"log"
	"mmacli/config/attack"
	"mmacli/config/defense"
	"mmacli/config/flags"
	"mmacli/models"
	"mmacli/myfmt"
)

var sampleNames = map[int]string{1: "Anderson Silva", 2: "Chael Sonnen"}

func CreateFighter(index int) models.Fighter {
	fmt.Printf("Type the name of fighter %d: ", index)
	name := ""
	if *UseMockFlag {
		name = sampleNames[index]
	} else {
		name = ReadString()
	}
	return models.Fighter{Name: name, Health: MaxHealth}
}

func SetupPlayers() {
	fighterOne := CreateFighter(1)
	fighterTwo := CreateFighter(2)

	WriteString(fmt.Sprintf("\n== %s versus %s ==\n\n", fighterOne.Name, fighterTwo.Name))
	TimeDelay()

	models.SetPlayers(map[int]*models.Fighter{1: &fighterOne, 2: &fighterTwo})
}

func SetupAttackList() {
	err := attack.LoadList("config/list.json")
	if err != nil {
		log.Fatal("Error on loading attack list:", err)
	}
}

func loadAssets() error {
	if err := flags.LoadList("config/flags/flags.json"); err != nil {
		return fmt.Errorf("Error loading feature flags: %e", err)
	}
	if err := attack.LoadList("config/attack/list.json"); err != nil {
		return fmt.Errorf("error on load attack list: %e", err)
	}
	if err := defense.LoadList("config/defense/list.json"); err != nil {
		return fmt.Errorf("error on load defense list: %e", err)
	}

	return nil
}

func setupCLIFlags() {
	flag.Parse()
	if *HideDelayFlag {
		Delay = 0
	}
}

func SetupGame() error {
	myfmt.PrintDelay("# Welcome to MMA CLI\n\n")
	setupCLIFlags()
	// TODO: Add Bruce Buffer introduce
	SetupPlayers()
	if err := loadAssets(); err != nil {
		return err
	}
	AnouncerPresentation()

	return nil
}
