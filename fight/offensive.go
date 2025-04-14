package fight

import (
	"fmt"
	"log"
	"mmacli/config/attack"
	"mmacli/config/flags"
	"mmacli/models"
	"mmacli/myfmt"
	"mmacli/utils"
	"strconv"
	"strings"
)

func ChooseAnAttack() (*attack.Attack, error) {
	myfmt.PrintDelay("Choose your attack play:\n")
	attackList := attack.GetAll()
	for _, attack := range attackList {
		bonus := getBonusString(attack.AttackBonus)
		fmt.Printf("[%d]: %s \t(Power: %s)\n", attack.Index, attack.Name, bonus)
	}
	choiceString, err := utils.ReadChar()
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

func getBonusString(bonus int) string {
	if bonus >= 0 {
		return fmt.Sprintf("+%d", bonus)
	}

	return fmt.Sprintf("%d", bonus)
}

func PlayerAttack(currentPlayerIndex int) error {
	players := models.GetPlayers()
	opponentIndex := utils.Swap[currentPlayerIndex]
	playerName := models.GetPlayerByIndex(currentPlayerIndex).Name
	opponentName := models.GetPlayerByIndex(opponentIndex).Name

	fmt.Printf("# %s (player %d) turn\n\n", playerName, currentPlayerIndex)

	attack, err := ChooseAnAttack()
	if err != nil {
		log.Fatal("Error on choose attack:", err)
	}

	// TODO: Clear the screen here but don't add time to clock
	myfmt.PrintDelay("\n- %s tries to hit a %s\n", playerName, strings.ToLower(attack.Name))
	power := utils.RollDice()
	myfmt.PrintDelay("- %s rolls: ", playerName)
	bonus := getBonusString(attack.AttackBonus)
	isCritical, isFail := handleCriticalAttack(power)
	criticalText := utils.GetCriticalText(isCritical, isFail)
	myfmt.PrintDelay("%d (%s) %s\n", power, bonus, criticalText)
	power += attack.AttackBonus

	if isFail {
		fmt.Printf("- %s missed the attack\n\n", playerName)
		return nil
	}

	if isCritical {
		bonus := utils.RollDice()
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
	opponent.Health -= damage
	if opponent.Health < 0 {
		opponent.Health = 0
	}

	return nil
}

func handleCriticalAttack(roll int) (isCritical bool, isFail bool) {
	enableCriticalHit := flags.GetById("enableCriticalHit")
	enableCriticalFail := flags.GetById("enableCriticalFail")

	isCritical = enableCriticalHit && roll == 6
	isFail = enableCriticalFail && roll == 1

	return
}
