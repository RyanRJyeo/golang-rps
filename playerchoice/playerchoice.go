package playerchoice

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/howeyc/gopass"
)

var weapons = []string{"Rock", "Paper", "Scissors"}

// ExposedChoice func
func ExposedChoice() string {
	var playerChoice string
	var input string
	fmt.Scanf("%s", &input)
	inputInt, _ := strconv.Atoi(input)

	playerPassed := false

	for playerPassed == false {
		if inputInt == 1 || inputInt == 2 || inputInt == 3 {
			playerPassed = true
			playerChoice = weapons[inputInt-1]
		} else {
			fmt.Print("Please type numbers 1, 2, or 3 to continue: ")
			fmt.Scanf("%s", &input)
			inputInt, _ = strconv.Atoi(input)
		}
	}

	return playerChoice
}

// HiddenChoice func
func HiddenChoice() string {
	var playerChoice string
	silentValue, _ := gopass.GetPasswdMasked()
	inputInt, _ := strconv.Atoi(string(silentValue))

	playerPassed := false

	for playerPassed == false {
		if inputInt == 1 || inputInt == 2 || inputInt == 3 {
			playerPassed = true
			playerChoice = weapons[inputInt-1]
		} else {
			fmt.Print("Please type numbers 1, 2, or 3 to continue: ")
			silentValue, _ = gopass.GetPasswdMasked()
			inputInt, _ = strconv.Atoi(string(silentValue))
		}
	}

	return playerChoice
}

// ComputerChoice func
func ComputerChoice() string {
	selector := rand.NewSource(time.Now().UnixNano())
	randomVal := rand.New(selector)
	return weapons[randomVal.Intn(3)]
}
