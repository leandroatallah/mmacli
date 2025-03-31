package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CheckGameOver(players map[int]*Fighter) bool {
	for _, p := range players {
		if p.health <= 0 {
			return true
		}
	}
	return false
}

func RollDice() int {
	return rand.Intn(6) + 1
}

func SetupPlayer(index int) Fighter {
	fmt.Printf("Type the name of fighter %d: ", index)
	name := ReadString()
	health := MaxHealth
	return Fighter{name, health}
}

func PlayersInitiative(players map[int]*Fighter) int {
	fmt.Println("Players roll initiative")
	time.Sleep(TurnDelay)
	for {
		playerOne := RollDice()
		playerTwo := RollDice()

		fmt.Print("Player one rolls: ")
		time.Sleep(TurnDelay)
		fmt.Println(playerOne)
		time.Sleep(TurnDelay)

		fmt.Print("Player two rolls: ")
		time.Sleep(TurnDelay)
		fmt.Println(playerTwo)
		time.Sleep(TurnDelay)

		if playerOne > playerTwo {
			return 1
		} else if playerTwo > playerOne {
			return 2
		}
		fmt.Println("Draw...")
		time.Sleep(TurnDelay)
	}
}

type Attack struct {
	name string
}

var attackList = map[string]Attack{
	"1": {"Jab"},
}

func PlayerAttack(currentPlayerIndex int, players map[int]*Fighter) error {
	opponentIndex := Swap[currentPlayerIndex]

	fmt.Println("Choose your play:")
	fmt.Println("[1]: Jab")
	choice, err := ReadChar()
	if err != nil {
		return err
	}

	attack, exists := attackList[choice]
	if !exists {
		return fmt.Errorf("Invalid attack choice")
	}
	switch attack {
	case attackList["1"]:
		fmt.Printf("Player %v tries to hit a punch\n", currentPlayerIndex)
		time.Sleep(TurnDelay)
		power := RollDice()
		fmt.Printf("Player %v rolls: ", currentPlayerIndex)
		time.Sleep(TurnDelay)
		fmt.Println(power)
		time.Sleep(TurnDelay)

		defense, err := PlayerDefense(opponentIndex, players)
		if err != nil {
			return fmt.Errorf("Oops! Something wrong happened.")
		}

		damage := max(power-defense, 0)
		if damage > 0 {
			fmt.Printf("Player %v suffered %d of damage\n", opponentIndex, damage)
		} else {
			fmt.Printf("Player %v doesn't suffered any damage\n", opponentIndex)
		}
		time.Sleep(TurnDelay)
		opponent := players[opponentIndex]
		opponent.health -= damage
		if opponent.health < 0 {
			opponent.health = 0
		}
	default:
		return fmt.Errorf("Oops! Something wrong happened.")
	}

	return nil
}

type Defense struct {
	name string
}

var defenseList = map[string]Defense{
	"1": {"Block"},
}

func PlayerDefense(defenderIndex int, players map[int]*Fighter) (int, error) {
	fmt.Printf("Player %v choose your defense:\n", defenderIndex)
	fmt.Println("[1]: Block")
	choice, err := ReadChar()
	if err != nil {
		return 0, err
	}

	defense, exists := defenseList[choice]
	if !exists {
		return 0, fmt.Errorf("Invalid defense choice")
	}

	time.Sleep(TurnDelay)

	switch defense {
	case defenseList["1"]:
		fmt.Printf("Player %d tries to block\n", defenderIndex)
		time.Sleep(TurnDelay)
		block := RollDice()
		fmt.Printf("Player %d rolls: ", defenderIndex)
		time.Sleep(TurnDelay)
		fmt.Println(block)
		return block, nil
	default:
		return 0, fmt.Errorf("Oops! Something wrong happened.")
	}
}
