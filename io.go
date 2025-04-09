package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadString() (text string) {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		text = scanner.Text()
	}
	return
}

func ReadChar() (string, error) {
	line := ReadString()
	if len(line) == 0 {
		return "", fmt.Errorf("No input given")
	}
	return string([]rune(line)[0]), nil

}

func WriteString(text string) error {
	out := os.Stdout
	_, err := out.WriteString(text)
	return err
}

func PrintStatus(players map[int]*Fighter) {
	for _, p := range players {
		fmt.Printf("\n%s\t%d\n", p.name, p.health)
	}
	fmt.Println()
}

func GetBonusString(bonus int) string {
	if bonus >= 0 {
		return fmt.Sprintf(" (+%d)", bonus)
	}

	return fmt.Sprintf(" (%d)", bonus)
}
