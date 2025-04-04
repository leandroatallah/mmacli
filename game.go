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

func GetPlayer(currentPlayerIndex int, players map[int]*Fighter) *Fighter {
	return players[currentPlayerIndex]
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

		p1Name := GetPlayer(1, players).name
		p2Name := GetPlayer(2, players).name
		fmt.Printf("%s rolls: ", p1Name)
		time.Sleep(TurnDelay)
		fmt.Println(playerOne)
		time.Sleep(TurnDelay)

		fmt.Printf("%s rolls: ", p2Name)
		time.Sleep(TurnDelay)
		fmt.Println(playerTwo)
		time.Sleep(TurnDelay)

		if playerOne > playerTwo {
			fmt.Printf("%s is next to play\n\n", p1Name)
			time.Sleep(TurnDelay)
			return 1
		} else if playerTwo > playerOne {
			fmt.Printf("%s is next to play\n\n", p2Name)
			time.Sleep(TurnDelay)
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
	playerName := GetPlayer(currentPlayerIndex, players).name
	opponentName := GetPlayer(opponentIndex, players).name

	fmt.Printf("%s (player %d) turn\n", playerName, currentPlayerIndex)

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
		fmt.Printf("%s tries to hit a punch\n", playerName)
		time.Sleep(TurnDelay)
		power := RollDice()
		isCritical := power == 6
		isFail := power == 1
		criticalText := ""
		if isCritical {
			criticalText = "[CRITICAL HIT!]"
		} else if isFail {
			criticalText = "[CRITICAL FAIL!]"
		}
		fmt.Printf("%s rolls: ", playerName)
		time.Sleep(TurnDelay)
		fmt.Printf("%d %s\n", power, criticalText)
		time.Sleep(TurnDelay)

		if isFail {
			fmt.Printf("%s missed the attack\n", playerName)
			return nil
		}

		if isCritical {
			bonus := RollDice()
			power += bonus
			fmt.Printf("%s rolls a bonus: ", playerName)
			time.Sleep(TurnDelay)
			fmt.Printf("%d\n", bonus)
			time.Sleep(TurnDelay)
		}

		defense, err := PlayerDefense(opponentIndex, players)
		if err != nil {
			return fmt.Errorf("Oops! Something wrong happened.")
		}

		damage := max(power-defense, 0)
		if damage > 0 {
			fmt.Printf("%s suffered %d of damage\n", opponentName, damage)
		} else {
			fmt.Printf("%s doesn't suffered any damage\n", opponentName)
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
	playerName := GetPlayer(defenderIndex, players).name
	fmt.Printf("%s choose your defense:\n", playerName)
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
		fmt.Printf("%s tries to block\n", playerName)
		time.Sleep(TurnDelay)
		block := RollDice()
		fmt.Printf("%s rolls: ", playerName)
		time.Sleep(TurnDelay)
		fmt.Println(block)
		return block, nil
	default:
		return 0, fmt.Errorf("Oops! Something wrong happened.")
	}
}
