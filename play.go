package main

import (
	"fmt"
	"log"
	"math"
	"mmacli/config/attack"
	"mmacli/config/defense"
	"mmacli/config/flags"
	"mmacli/myfmt"
	"strconv"
	"strings"
)

func PlayersInitiative(players map[int]*Fighter) int {
	myfmt.PrintDelay("\n# Players roll initiative\n\n")
	for {
		type p struct {
			name string
			roll int
		}
		p1 := p{GetPlayer(1, players).name, RollDice()}
		p2 := p{GetPlayer(2, players).name, RollDice()}

		for _, p := range []p{p1, p2} {
			myfmt.PrintDelay("- %s rolls: ", p.name)
			myfmt.PrintDelay("%d", p.roll)
		}

		if p1.roll > p2.roll {
			myfmt.PrintDelay("- %s (player 1) is next to play\n\n", p1.name)
			return 1
		} else if p2.roll > p1.roll {
			myfmt.PrintDelay("- %s (player 2) is next to play\n\n", p2.name)
			return 2
		}
		myfmt.PrintDelay("- Draw...")
	}
}

func chooseAnAttack() (*attack.Attack, error) {
	myfmt.PrintDelay("Choose your attack play:")
	attackList := attack.GetAll()
	for index, attack := range attackList {
		bonus := GetBonusString(attack.AttackBonus)
		fmt.Printf("[%d]: %s \t(Power: %s)\n", index, attack.Name, bonus)
	}
	choiceString, err := ReadChar()
	if err != nil {
		return nil, err
	}

	choice, err := strconv.Atoi(choiceString)
	if err != nil {
		return nil, err
	}

	attack, exists := attack.GetByIndex(choice)
	if !exists {
		// TODO: Handle this error
		return nil, fmt.Errorf("Invalid attack choice")
	}

	return attack, nil
}

func handleCriticalAttack(roll int) (isCritical bool, isFail bool) {
	enableCriticalHit := flags.GetById("enableCriticalHit")
	enableCriticalFail := flags.GetById("enableCriticalFail")

	isCritical = enableCriticalHit && roll == 6
	isFail = enableCriticalFail && roll == 1

	return
}

func getCriticalAttackText(isCritical, isFail bool) string {
	if isCritical {
		return "[CRITICAL HIT!]"
	} else if isFail {
		return "[CRITICAL FAIL!]"
	}

	return ""
}

func PlayerAttack(currentPlayerIndex int, players map[int]*Fighter) error {
	opponentIndex := Swap[currentPlayerIndex]
	playerName := GetPlayer(currentPlayerIndex, players).name
	opponentName := GetPlayer(opponentIndex, players).name

	fmt.Printf("%s (player %d) turn\n", playerName, currentPlayerIndex)

	attack, err := chooseAnAttack()
	if err != nil {
		log.Fatal("Error on choose attack:", err)
	}

	myfmt.PrintDelay("\n- %s tries to hit a %s\n", playerName, strings.ToLower(attack.Name))
	power := RollDice()
	myfmt.PrintDelay("- %s rolls: ", playerName)
	bonus := GetBonusString(attack.AttackBonus)
	isCritical, isFail := handleCriticalAttack(power)
	criticalText := getCriticalAttackText(isCritical, isFail)
	myfmt.PrintDelay("%d (%s) %s\n", power, bonus, criticalText)
	power += attack.AttackBonus

	if isFail {
		fmt.Printf("- %s missed the attack\n\n", playerName)
		return nil
	}

	if isCritical {
		bonus := RollDice()
		power += bonus
		myfmt.PrintDelay("- %s rolls a bonus: ", playerName)
		myfmt.PrintDelay("%d\n", bonus)
		myfmt.PrintDelay("- %s attack is: %d (%d + %d)\n\n", playerName, power, power-bonus, bonus)
	}

	defense, err := PlayerDefense(opponentIndex, attack.AttackBonus, players)
	if err != nil {
		return fmt.Errorf("Oops! Something wrong happened.")
	}

	damage := max(power-defense, 0)
	if damage > 0 {
		myfmt.PrintDelay("- %s suffered %d of damage\n", opponentName, damage)
	} else {
		myfmt.PrintDelay("- %s doesn't suffered any damage\n", opponentName)
	}
	opponent := players[opponentIndex]
	opponent.health -= damage
	if opponent.health < 0 {
		opponent.health = 0
	}

	return nil
}

type Defense struct {
	name     string
	accuracy int // Minimum roll required for success
}

func chooseAnDefense(playerName string, bonus int) (*defense.Defense, error) {
	fmt.Printf("\n%s choose your defense:\n", playerName)
	defenseList := defense.GetAll()
	for index, defense := range defenseList {
		// TODO: Improve varible names
		accuracy := defense.Accuracy
		if defense.Accuracy != 6 {
			accuracy = min(5, accuracy+bonus)
		}
		successChance := float64(MaxDiceNumber - accuracy)
		precision := 1 - (successChance / float64(MaxDiceNumber))
		precisionFormated := min(100, int((precision)*100))
		fmt.Printf("[%d]: %s \t(%d%%)\n", index, defense.Name, precisionFormated)
	}
	choiceString, err := ReadChar()
	if err != nil {
		return nil, err
	}

	choice, err := strconv.Atoi(choiceString)
	if err != nil {
		return nil, err
	}

	defense, exists := defense.GetByIndex(choice)
	if !exists {
		return nil, fmt.Errorf("Invalid defense choice")
	}

	return defense, nil
}

func PlayerDefense(defenderIndex, bonus int, players map[int]*Fighter) (int, error) {
	playerName := GetPlayer(defenderIndex, players).name
	defense, err := chooseAnDefense(playerName, bonus)
	if err != nil {
		return 0, err
	}

	myfmt.PrintDelay("")
	defenseName := strings.ToLower(defense.Name)
	myfmt.PrintDelay("\n- %s tries to %s\n", playerName, defenseName)
	roll := RollDice()
	myfmt.PrintDelay("- %s rolls: ", playerName)
	myfmt.PrintDelay("%d\n", roll)

	switch defense.Index {
	case 1:
		return roll, nil
	case 2, 3:
		if defense.Accuracy+bonus > MaxDiceNumber-roll {
			fmt.Printf("- %s successfully %ss\n", playerName, defenseName)
			return math.MaxInt, nil
		}
		fmt.Printf("- %s fails to %s\n", playerName, defenseName)
		return 0, nil
	}

	return 0, fmt.Errorf("Oops! Something wrong happened.")
}
