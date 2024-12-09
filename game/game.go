package game

import (
	"bufio"
	"fmt"
	"os"
)

// 4 одно-палубных
// 3 двух-палубных
// 2 трёх-палубных
// 1 четырёх-палубный

const BOARD_SIZE int = 10

type GameBoard = [BOARD_SIZE][BOARD_SIZE]bool

func GetEmptyBoard() GameBoard {
	var board GameBoard = GameBoard{}

	for i := 0; i < BOARD_SIZE; i++ {
		row := [BOARD_SIZE]bool{}

		for j := 0; j < BOARD_SIZE; j++ {
			row[j] = false
		}

		board[i] = row
	}

	return board
}

//func randomizeShips

var rowsAlphabet = [BOARD_SIZE]string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
}

func PrintGameBoard(board GameBoard) {
	fmt.Print("   ")

	for i := range rowsAlphabet {
		fmt.Printf("%v  ", rowsAlphabet[i])
	}

	for i, row := range board {
		//fmt.Printf("\n%v", row)
		fmt.Println()

		if i == 9 {
			fmt.Printf("%v ", i+1)
		} else {
			fmt.Printf("%v  ", i+1)
		}

		for range row {
			fmt.Printf("%v  ", "◻") //🔥❌◻
		}
	}
}

func Turn() {

}

func getUserInput(msg string) (string, error) {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	result, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Println("Something went wrong occure you input a turn. Please try again.")
	}

}
