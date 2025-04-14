package utils

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

func GetCriticalText(isCritical, isFail bool) string {
	if isCritical {
		return "[CRITICAL HIT!]"
	} else if isFail {
		return "[CRITICAL FAIL!]"
	}

	return ""
}
