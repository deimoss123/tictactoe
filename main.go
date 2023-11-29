package main

import (
	"fmt"
	"os"
	"tictactoe/game"
)

func main() {
	fmt.Print("\033[H\033[2J")
	// message := "_   _      _             _             \n" +
	// 	"| |_(_) ___| |_ __ _  ___| |_ ___   ___ \n" +
	// 	"| __| |/ __| __/ _` |/ __| __/ _ \\ / _ \\\n" +
	// 	"| |_| | (__| || (_| | (__| || (_) |  __/\n" +
	// 	" \\__|_|\\___|\\__\\__,_|\\___|\\__\\___/ \\___|\n"

	message := `
 ____           _ _        _ 
/ ___|_   _____(_) | _____| |
\___ \ \ / / _ \ | |/ / __| |
 ___) \ V /  __/ |   <\__ \_|
|____/ \_/ \___|_|_|\_\___(_)
`

	fmt.Println(message)
	fmt.Print("Spied 'Enter' lai uzsāktu spēli... ")
	var input string
	fmt.Scanln(&input)

	if input == "exit" {
		os.Exit(0)
	}

	hasExited := false

	for !hasExited {
		currGame := game.New()

		for game.IsActive(&currGame) {
			game.PrintBoard(&currGame)
			num := game.GetNextMoveInput(&currGame)
			game.UpdateGame(&currGame, num)
		}

		game.PrintBoard(&currGame)

		fmt.Print("Spēle beidzās. Spied 'Enter' lai spēlētu vēlreiz... ")

		fmt.Scanln(&input)

		if input == "exit" {
			os.Exit(0)
		}
	}
}
