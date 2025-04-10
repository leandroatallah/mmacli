package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[2J\033[H")
	}
}

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
	ClearScreen()

	clock := GetClockTime()

	p1NameLen := len(players[1].name)
	p2NameLen := len(players[2].name)
	damageOffset := 6
	playersNameLen := max(p1NameLen, p2NameLen, 9) + damageOffset

	// line
	line := ""
	for range playersNameLen + 8 {
		line += "#"
	}
	line += "\n"

	// line with spaces
	lineWithSpaces := "## "
	for range playersNameLen + 2 {
		lineWithSpaces += " "
	}
	lineWithSpaces += " ##"

	// Clock
	lineClock := func() string {
		leading := ""
		additionalSpace := playersNameLen%2 == 0
		for range (playersNameLen - 7) / 2 {
			leading += " "
		}
		clockTrim := fmt.Sprintf("[%v]", strings.TrimSpace(clock))
		result := "##  " + leading + clockTrim + leading
		if additionalSpace {
			result += " "
		}
		result += "  ##"

		return result
	}

	// Round
	lineRound := func() string {
		leading := ""
		additionalSpace := playersNameLen%2 == 0
		for range (playersNameLen - 7) / 2 {
			leading += " "
		}
		result := fmt.Sprintf("##  %sRound %s%s", leading, GetRound(), leading)
		if additionalSpace {
			result += " "
		}
		result += "  ##"
		return result
	}

	// Players name
	playersName := func() string {
		result := ""
		for _, p := range players {
			remainSpace := playersNameLen - len(p.name) - 2
			result += fmt.Sprintf("##  %s", p.name)
			for range remainSpace {
				result += " "
			}
			if p.health < 10 {
				result += " "
			}
			result += fmt.Sprintf("%d  ##\n", p.health)
		}
		return result
	}

	// Print status
	fmt.Print(line)
	fmt.Println(lineWithSpaces)
	fmt.Println(lineRound())
	fmt.Println(lineClock())
	fmt.Println(lineWithSpaces)
	fmt.Print(playersName())
	fmt.Println(lineWithSpaces)
	fmt.Println(line)
}

func GetBonusString(bonus int) string {
	if bonus >= 0 {
		return fmt.Sprintf("+%d", bonus)
	}

	return fmt.Sprintf("%d", bonus)
}
