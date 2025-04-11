package models

type Fighter struct {
	Name   string
	Health int
}

var players map[int]*Fighter

func GetPlayers() map[int]*Fighter {
	return players
}

func GetPlayerByIndex(currentPlayerIndex int) *Fighter {
	return players[currentPlayerIndex]
}

func SetPlayers(value map[int]*Fighter) {
	players = value
}
