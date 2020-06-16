package gamemodes

import (
	"fmt"

	"github.com/RyanRJyeo/go_crash_course/16_rps/gamelogic"
	"github.com/RyanRJyeo/go_crash_course/16_rps/playerchoice"
	"github.com/RyanRJyeo/go_crash_course/16_rps/speech"
)

// PlayMode func
func PlayMode() {
	speech.SelectedMode("Play")
	speech.ChooseWeapon("Player")

	player1 := "Player"
	player2 := "Computer"

	player1Choice := playerchoice.ExposedChoice()
	player2Choice := playerchoice.ComputerChoice()

	gamelogic.EndGame(player1, player2, player1Choice, player2Choice)
}

// WatchMode func
func WatchMode() {
	speech.SelectedMode("Watch")

	player1 := "Computer 1"
	player2 := "Computer 2"

	player1Choice := playerchoice.ComputerChoice()
	player2Choice := playerchoice.ComputerChoice()

	gamelogic.EndGame(player1, player2, player1Choice, player2Choice)
}

// VersusMode func
func VersusMode() {
	// Player 1 starts
	speech.SelectedMode("Versus")
	speech.ChooseWeapon("Player 1")

	player1 := "Player 1"
	player1Choice := playerchoice.HiddenChoice()

	// Player 2 continues
	fmt.Print("\n=======================================\n\n")
	speech.ChooseWeapon("Player 2")

	player2 := "Player 2"
	player2Choice := playerchoice.HiddenChoice()

	gamelogic.EndGame(player1, player2, player1Choice, player2Choice)
}
