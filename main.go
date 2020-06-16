package main

import (
	"fmt"

	"github.com/RyanRJyeo/go_crash_course/16_rps/gamemodes"
	"github.com/RyanRJyeo/go_crash_course/16_rps/speech"
)

func main() {

	// Game welcome message and instructions
	speech.Welcome()
	speech.ChooseGameMode()
	var input string
	fmt.Scanf("%s", &input)
	gameStarted := true

	for gameStarted == true {

		if input == "1" {

			// Player choose PlayMode
			gamemodes.PlayMode()
			gameStarted = false

		} else if input == "2" {

			// Player choose WatchMode
			gamemodes.WatchMode()
			gameStarted = false

		} else if input == "3" {

			// Player choose VersusMode
			gamemodes.VersusMode()
			gameStarted = false

		} else if input == "4" {

			// Player choose to quite the game
			fmt.Print("\n\nBye Bye!\n\n")
			return

		} else {

			// Error message if player don't type 1, 2, 3, or 4
			speech.Error()
			fmt.Scanf("%s", &input)

		}

	}

	// Option to replay the game
	speech.Replay()
	var replay string
	fmt.Scanf("%s", &replay)
	if replay == "1" {
		main()
	} else {
		fmt.Print("\n\nBye Bye!\n\n")
		return
	}
}
