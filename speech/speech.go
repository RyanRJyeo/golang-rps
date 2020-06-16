package speech

import "fmt"

// Welcome message
func Welcome() {
	fmt.Println("\033[2J")
	fmt.Print("\n=======================================\n\n")
	fmt.Println("Welcome to The RPS Game!!!")
	fmt.Println("\nFollowing are the game modes that you can choose from:")
	fmt.Println("\nPlay mode: 	Play against the computer")
	fmt.Println("Watch mode: 	Watch 2 computers play against each other")
	fmt.Println("Versus mode: 	Play against another player")
}

// ChooseGameMode message
func ChooseGameMode() {
	fmt.Println("\nPlease type the following numbers to continue:")
	fmt.Print("\n-------------------------------------------------------")
	fmt.Print("\n1 - Play mode \n2 - Watch mode \n3 - Versus mode \n4 - Quit game\n")
	fmt.Print("-------------------------------------------------------\n\n")
	fmt.Print("Enter your choice: ")
}

// SelectedMode message
func SelectedMode(game string) {
	fmt.Println("\033[2J")
	fmt.Print("\n=======================================\n\n")
	fmt.Println("You have chosen " + game + " Mode! ")
}

// ChooseWeapon message
func ChooseWeapon(player string) {
	fmt.Println(player + " please type the following numbers to continue:")
	fmt.Print("\n-------------------------------------------------------")
	fmt.Print("\n1 - Rock\n2 - Paper\n3 - Scissors\n")
	fmt.Print("-------------------------------------------------------\n\n")
	fmt.Print("Enter your choice: ")
}

// WeaponsChosen message
func WeaponsChosen(player1, player2, player1Choice, player2Choice string) {
	fmt.Print("\n=======================================\n")
	fmt.Println(player1 + " choice is: " + player1Choice)
	fmt.Println(player2 + " choice is: " + player2Choice)
	fmt.Print("=======================================\n\n")
}

// Error message
func Error() {
	fmt.Println("\033[2J")
	fmt.Print("\n=======================================\n\n")
	fmt.Println("Sorry that is an invalid input")
	fmt.Println("Please type the following numbers to continue:")
	fmt.Print("\n-------------------------------------------------------")
	fmt.Print("\n1 - Play mode \n2 - Watch mode \n3 - Versus mode \n4 - Quit game\n")
	fmt.Print("-------------------------------------------------------\n\n")
	fmt.Print("Enter your choice: ")
}

// Replay message
func Replay() {
	fmt.Println("Would you like to play again?")
	fmt.Print("\n-------------------------------------------------------")
	fmt.Print("\nType '1' to play again \nType anything else to quit the game\n")
	fmt.Print("-------------------------------------------------------\n\n")
	fmt.Print("Enter your choice: ")
}
