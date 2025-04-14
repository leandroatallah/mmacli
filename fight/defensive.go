package fight

import (
	"fmt"
	"math"
	"mmacli/config/defense"
	"mmacli/models"
	"mmacli/myfmt"
	"mmacli/utils"
	"strconv"
	"strings"
)

var defenseList []defense.Defense

func ChooseAnDefense(playerName string, bonus int) (*defense.Defense, error) {
	fmt.Printf("\n%s choose your defense:\n", playerName)
	defenseList := defense.GetAll()
	for _, defense := range defenseList {
		accuracy := defense.Accuracy
		if defense.Accuracy != 6 {
			accuracy = min(5, accuracy+bonus)
		}
		successChance := float64(6 - accuracy)
		precision := 1 - (successChance / float64(6))
		precisionFormated := min(100, int((precision)*100))
		fmt.Printf("[%d]: %s \t(%d%%)\n", defense.Index, defense.Name, precisionFormated)
	}
	choiceString, err := utils.ReadChar()
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

func PlayerDefense(defenderIndex, bonus int, players map[int]*models.Fighter) (int, error) {
	playerName := models.GetPlayerByIndex(defenderIndex).Name
	defense, err := ChooseAnDefense(playerName, bonus)
	if err != nil {
		return 0, err
	}

	myfmt.PrintDelay("")
	defenseName := strings.ToLower(defense.Name)
	myfmt.PrintDelay("\n- %s tries to %s\n", playerName, defenseName)
	roll := utils.RollDice()
	myfmt.PrintDelay("- %s rolls: ", playerName)
	myfmt.PrintDelay("%d\n", roll)

	switch defense.Index {
	case 1:
		return roll, nil
	case 2, 3:
		if defense.Accuracy+bonus > 6-roll {
			fmt.Printf("- %s successfully %ss\n", playerName, defenseName)
			return math.MaxInt, nil
		}
		fmt.Printf("- %s fails to %s\n", playerName, defenseName)
		return 0, nil
	}

	return 0, fmt.Errorf("Oops! Something wrong happened.")
}
