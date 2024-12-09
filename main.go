package main

import (
	"fmt"
	"sea_battle/game"
	"sea_battle/helpers"
)

func main() {
	fmt.Println("Welcome to the sea battle")
	fmt.Println("Rules of this game is very easy:")
	fmt.Println("1) You need fire on board peases with writing duplex coords, like A1, C8, B4 in console")
	fmt.Println("2) You need kill all enemy ships")
	fmt.Println("3) If you get on a ship, board pease will change symbol on 🔥")
	fmt.Println("3) If you kill a ship, board pease will change symbol on ❌")
	fmt.Println("Good luck")

	board := game.GetBoard()

	for {
		game.PrintGameBoard(board)

		inputText := helpers.GetUserInput("Please write cords")

		cords, err := helpers.ParseUserInputToCords(inputText)

		if err != nil {
			fmt.Printf("Input error: %v", err)
			continue
		}

		game.Turn(&board, *cords)
	}
}
