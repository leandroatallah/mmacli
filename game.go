package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
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
	fmt.Printf("\n# Players roll initiative\n\n")
	time.Sleep(Delay)
	for {
		playerOne := RollDice()
		playerTwo := RollDice()

		p1Name := GetPlayer(1, players).name
		p2Name := GetPlayer(2, players).name
		fmt.Printf("- %s rolls: ", p1Name)
		time.Sleep(Delay)
		fmt.Println(playerOne)
		time.Sleep(Delay)

		fmt.Printf("- %s rolls: ", p2Name)
		time.Sleep(Delay)
		fmt.Println(playerTwo)
		time.Sleep(Delay)

		if playerOne > playerTwo {
			fmt.Printf("- %s (player 1) is next to play\n\n", p1Name)
			time.Sleep(Delay)
			return 1
		} else if playerTwo > playerOne {
			fmt.Printf("- %s (player 2) is next to play\n\n", p2Name)
			time.Sleep(Delay)
			return 2
		}
		fmt.Println("- Draw...")
		time.Sleep(Delay)
	}
}

type Attack struct {
	name  string
	bonus int
	// accuracy
}

var attackList = map[string]Attack{
	"1": {"Jab", 0},
	"2": {"Cross", 1},
	"3": {"Uppercut", 2},
	// "4": {"Hook"},
}

func PlayerAttack(currentPlayerIndex int, players map[int]*Fighter) error {
	opponentIndex := Swap[currentPlayerIndex]
	playerName := GetPlayer(currentPlayerIndex, players).name
	opponentName := GetPlayer(opponentIndex, players).name

	fmt.Printf("%s (player %d) turn\n", playerName, currentPlayerIndex)

	fmt.Println("Choose your attack play:")
	for index, attack := range attackList {
		fmt.Printf("[%s]: %s \t(Power: +%d)\n", index, attack.name, attack.bonus)
	}
	choice, err := ReadChar()
	if err != nil {
		return err
	}

	attack, exists := attackList[choice]
	if !exists {
		return fmt.Errorf("Invalid attack choice")
	}
	fmt.Printf("\n- %s tries to hit a %s\n", playerName, strings.ToLower(attack.name))
	time.Sleep(Delay)

	power := RollDice()
	isCritical := power == 6
	isFail := power == 1
	criticalText := ""
	if isCritical {
		criticalText = "[CRITICAL HIT!]"
	} else if isFail {
		criticalText = "[CRITICAL FAIL!]"
	}
	fmt.Printf("- %s rolls: ", playerName)
	time.Sleep(Delay)
	fmt.Printf("%d (+%d) %s\n", power, attack.bonus, criticalText)
	power += attack.bonus
	time.Sleep(Delay)

	if isFail {
		fmt.Printf("- %s missed the attack\n\n", playerName)
		return nil
	}

	if isCritical {
		bonus := RollDice()
		power += bonus
		fmt.Printf("- %s rolls a bonus: ", playerName)
		time.Sleep(Delay)
		fmt.Printf("%d\n", bonus)
		time.Sleep(Delay)
		fmt.Printf("- %s attack is: %d (%d + %d)\n\n", playerName, power, power-bonus, bonus)
		time.Sleep(Delay)
	}

	defense, err := PlayerDefense(opponentIndex, players)
	if err != nil {
		return fmt.Errorf("Oops! Something wrong happened.")
	}

	damage := max(power-defense, 0)
	if damage > 0 {
		fmt.Printf("- %s suffered %d of damage\n", opponentName, damage)
	} else {
		fmt.Printf("- %s doesn't suffered any damage\n", opponentName)
	}
	time.Sleep(Delay)
	opponent := players[opponentIndex]
	opponent.health -= damage
	if opponent.health < 0 {
		opponent.health = 0
	}

	return nil
}

type Defense struct {
	name      string
	precision string
}

var defenseList = map[string]Defense{
	"1": {"Block", "100%"},
	"2": {"Evade", "66%"},
	"3": {"Counter", "50%"},
}

func PlayerDefense(defenderIndex int, players map[int]*Fighter) (int, error) {
	playerName := GetPlayer(defenderIndex, players).name
	fmt.Printf("\n%s choose your defense:\n", playerName)
	for index, defense := range defenseList {
		fmt.Printf("[%s]: %s \t(%s)\n", index, defense.name, defense.precision)
	}
	choice, err := ReadChar()
	if err != nil {
		return 0, err
	}

	defense, exists := defenseList[choice]
	if !exists {
		return 0, fmt.Errorf("Invalid defense choice")
	}

	defenseName := strings.ToLower(defense.name)

	time.Sleep(Delay)
	fmt.Printf("\n- %s tries to %s\n", playerName, defenseName)
	time.Sleep(Delay)
	roll := RollDice()
	fmt.Printf("- %s rolls: ", playerName)
	time.Sleep(Delay)
	fmt.Printf("%d\n", roll)
	time.Sleep(Delay)
	switch defense {
	case defenseList["1"]:
		return roll, nil
	case defenseList["2"]:
		result := 0
		if roll > 2 {
			fmt.Printf("- %s succesfully %ss\n", playerName, defenseName)
			result = math.MaxInt
		} else {
			fmt.Printf("- %s fails to %s\n", playerName, defenseName)
		}
		return result, nil
	case defenseList["3"]:
		result := 0
		if roll > 3 {
			fmt.Printf("- %s succesfully %ss\n", playerName, defenseName)
			result = math.MaxInt
		} else {
			fmt.Printf("- %s fails to %s\n", playerName, defenseName)
		}
		return result, nil
	}

	return 0, fmt.Errorf("Oops! Something wrong happened.")

}
