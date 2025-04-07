package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"mmacli/config"
	"strconv"
	"strings"
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

func PlayersInitiative(players map[int]*Fighter) int {
	fmt.Printf("\n# Players roll initiative\n\n")
	TimeDelay()
	for {
		playerOne := RollDice()
		playerTwo := RollDice()

		p1Name := GetPlayer(1, players).name
		p2Name := GetPlayer(2, players).name
		fmt.Printf("- %s rolls: ", p1Name)
		TimeDelay()
		fmt.Println(playerOne)
		TimeDelay()

		fmt.Printf("- %s rolls: ", p2Name)
		TimeDelay()
		fmt.Println(playerTwo)
		TimeDelay()

		if playerOne > playerTwo {
			fmt.Printf("- %s (player 1) is next to play\n\n", p1Name)
			TimeDelay()
			return 1
		} else if playerTwo > playerOne {
			fmt.Printf("- %s (player 2) is next to play\n\n", p2Name)
			TimeDelay()
			return 2
		}
		fmt.Println("- Draw...")
		TimeDelay()
	}
}

func chooseAnAttack() (*config.Attack, error) {
	fmt.Println("Choose your attack play:")
	attackList := config.GetAllAttackList()
	for index, attack := range attackList {
		plusSign := ""
		if attack.AttackBonus >= 0 {
			plusSign = "+"
		}
		fmt.Printf("[%d]: %s \t(Power: %s%d)\n", index, attack.Name, plusSign, attack.AttackBonus)
	}
	choiceString, err := ReadChar()
	if err != nil {
		return nil, err
	}

	choice, err := strconv.Atoi(choiceString)
	if err != nil {
		return nil, err
	}

	attack, exists := config.GetAttackByIndex(choice)
	if !exists {
		// TODO: Handle this error
		return nil, fmt.Errorf("Invalid attack choice")
	}

	return attack, nil
}

func PlayerAttack(currentPlayerIndex int, players map[int]*Fighter) error {
	err := config.LoadAttackList("config/attacks.json")
	if err != nil {
		log.Fatal("error on load attack list:", err)
	}

	opponentIndex := Swap[currentPlayerIndex]
	playerName := GetPlayer(currentPlayerIndex, players).name
	opponentName := GetPlayer(opponentIndex, players).name

	fmt.Printf("%s (player %d) turn\n", playerName, currentPlayerIndex)

	attack, err := chooseAnAttack()
	if err != nil {
		log.Fatal("Error on choose attack:", err)
	}

	fmt.Printf("\n- %s tries to hit a %s\n", playerName, strings.ToLower(attack.Name))
	TimeDelay()

	power := RollDice()
	isCritical := config.Flags["enableCriticalHit"] && power == 6
	isFail := config.Flags["enableCriticalFail"] && power == 1
	criticalText := ""
	if isCritical {
		criticalText = "[CRITICAL HIT!]"
	} else if isFail {
		criticalText = "[CRITICAL FAIL!]"
	}
	fmt.Printf("- %s rolls: ", playerName)
	TimeDelay()
	fmt.Printf("%d (+%d) %s\n", power, attack.AttackBonus, criticalText)
	power += attack.AttackBonus
	TimeDelay()

	if isFail {
		fmt.Printf("- %s missed the attack\n\n", playerName)
		return nil
	}

	if isCritical {
		bonus := RollDice()
		power += bonus
		fmt.Printf("- %s rolls a bonus: ", playerName)
		TimeDelay()
		fmt.Printf("%d\n", bonus)
		TimeDelay()
		fmt.Printf("- %s attack is: %d (%d + %d)\n\n", playerName, power, power-bonus, bonus)
		TimeDelay()
	}

	defense, err := PlayerDefense(opponentIndex, attack.AttackBonus, players)
	if err != nil {
		return fmt.Errorf("Oops! Something wrong happened.")
	}

	damage := max(power-defense, 0)
	if damage > 0 {
		fmt.Printf("- %s suffered %d of damage\n", opponentName, damage)
	} else {
		fmt.Printf("- %s doesn't suffered any damage\n", opponentName)
	}
	TimeDelay()
	opponent := players[opponentIndex]
	opponent.health -= damage
	if opponent.health < 0 {
		opponent.health = 0
	}

	return nil
}

type Defense struct {
	name      string
	threshold int // Minimum roll required for success
}

// TODO: Convert to a JSON file
var defenseList = map[string]Defense{
	"1": {"Block", 6},   // 100%
	"2": {"Evade", 4},   // 66%
	"3": {"Counter", 3}, // 50%
}

func PlayerDefense(defenderIndex, bonus int, players map[int]*Fighter) (int, error) {
	playerName := GetPlayer(defenderIndex, players).name
	fmt.Printf("\n%s choose your defense:\n", playerName)
	for index, defense := range defenseList {
		// TODO: Improve varible names
		threshold := defense.threshold
		if defense.threshold != 6 {
			threshold = min(5, threshold+bonus)
		}
		successChance := float64(MaxDiceNumber - threshold)
		precision := 1 - (successChance / float64(MaxDiceNumber))
		precisionFormated := min(100, int((precision)*100))
		fmt.Printf("[%s]: %s \t(%d%%)\n", index, defense.name, precisionFormated)
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

	TimeDelay()
	fmt.Printf("\n- %s tries to %s\n", playerName, defenseName)
	TimeDelay()
	roll := RollDice()
	fmt.Printf("- %s rolls: ", playerName)
	TimeDelay()
	fmt.Printf("%d\n", roll)
	TimeDelay()
	switch defense {
	case defenseList["1"]:
		return roll, nil
	case defenseList["2"], defenseList["3"]:
		if defense.threshold+bonus > MaxDiceNumber-roll {
			fmt.Printf("- %s successfully %ss\n", playerName, defenseName)
			return math.MaxInt, nil
		}
		fmt.Printf("- %s fails to %s\n", playerName, defenseName)
		return 0, nil
	}

	return 0, fmt.Errorf("Oops! Something wrong happened.")

}
