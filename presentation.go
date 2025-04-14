package main

import (
	"fmt"
	"mmacli/cliflags"
	"mmacli/models"
	"mmacli/myfmt"
	"mmacli/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var delay = 1 * time.Second

func TimeDelay() {
	time.Sleep(delay)
}

func SetTimeDelay(value time.Duration) {
	delay = value
}

func ClearScreen() {
	if cliflags.GetCliFlag("preserve") || cliflags.GetCliFlag("p") {
		return
	}

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[2J\033[H")
	}
}

func PressEnter() {
	fmt.Printf("\n\nPress [ENTER] to continue.")
	if cliflags.GetCliFlag("noenter") || cliflags.GetCliFlag("n") {
		return
	}
	utils.ReadChar()
}

func PrintStatus() {
	ClearScreen()
	players := models.GetPlayers()
	clock := GetClockTime()

	// Get names length. Minimun of 9 plus 6 offset
	nameLen := max(len(players[1].Name), len(players[2].Name), 9) + 6
	padding := strings.Repeat(" ", (nameLen-6)/2)

	line := strings.Repeat("#", nameLen+8) + "\n"
	lineWithSpaces := "## " + strings.Repeat(" ", nameLen+2) + " ##"

	lineClock := func() string {
		clockStr := fmt.Sprintf("[%v]", strings.TrimSpace(clock))
		extra := ""
		if nameLen%2 == 0 {
			extra = " "
		}
		return "## " + padding + clockStr + padding + extra + " ##"
	}

	lineRound := "## " + padding + "Round " + GetRound() + padding + "  ##"

	playersName := func() string {
		result := ""
		for _, p := range players {
			result += fmt.Sprintf("##  %s", p.Name)
			result += strings.Repeat(" ", nameLen-len(p.Name)-2)
			if p.Health < 10 {
				result += " "
			}
			result += fmt.Sprintf("%d  ##\n", p.Health)
		}
		return result
	}

	fmt.Print(line)
	fmt.Println(lineWithSpaces)
	fmt.Println(lineRound)
	fmt.Println(lineClock())
	fmt.Println(lineWithSpaces)
	fmt.Print(playersName())
	fmt.Println(lineWithSpaces)
	fmt.Println(line)
}

func AnouncerPresentation() {
	players := models.GetPlayers()

	myfmt.PrintDelay("- \"Ladies and Gentlemen.\"\n")
	// myfmt.PrintDelay("- \"This is the moment you’ve all been waiting for!\"\n")
	myfmt.PrintDelay("- \"This is the main event of the evening!\"\n")
	myfmt.PrintDelay("- \"And now... ")
	myfmt.PrintDelay("It's time!\"\n")
	myfmt.PrintDelay("- \"Introducing first...\"\n")
	myfmt.PrintDelay("- \"In the left corner: ")
	myfmt.PrintDelay("%s\"\n", players[1].Name)
	myfmt.PrintDelay("- \"And now, introducing his opponent...\"\n")
	myfmt.PrintDelay("- \"In the right corner: ")
	myfmt.PrintDelay("%s\"\n", players[2].Name)
}
