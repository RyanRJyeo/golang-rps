package gamelogic

import (
	"fmt"

	"github.com/RyanRJyeo/go_crash_course/16_rps/speech"
)

// EndGame func
func EndGame(player1, player2, player1Choice, player2Choice string) {

	// Display players' choices
	speech.WeaponsChosen(player1, player2, player1Choice, player2Choice)

	// Define win conditions
	var winConditions = map[string][]string{
		"Rock":     {"Scissors"},
		"Scissors": {"Paper"},
		"Paper":    {"Rock"},
	}

	// Check who wins the game
	b := contains(winConditions[player1Choice], player2Choice)
	if b == true {
		fmt.Println(player1 + " is the winner!")
	} else if player1Choice == player2Choice {
		fmt.Println("It's a draw!")
	} else {
		fmt.Println(player2 + " is the winner!")
	}
}

// Check if array contains the string or not
func contains(arr []string, str string) bool {
	for _, value := range arr {
		if value == str {
			return true
		}
	}
	return false
}
