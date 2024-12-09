package main

import (
	"sea_battle/game"
)

func main() {
	board := game.GetEmptyBoard()

	game.PrintGameBoard(board)
}
