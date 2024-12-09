package game

import (
	"errors"
	"fmt"
)

// 4 одно-палубных
// 3 двух-палубных
// 2 трёх-палубных
// 1 четырёх-палубный

const BOARD_SIZE int = 10

type ShipDeck struct {
	isTouched bool
	isDead    bool
}

type GameBoard = [BOARD_SIZE][BOARD_SIZE]*ShipDeck

func GetBoard() GameBoard {
	board := getEmptyBoard()

	RandomizeShips(&board)

	return board
}

func getEmptyBoard() GameBoard {
	var board GameBoard = GameBoard{}

	for i := 0; i < BOARD_SIZE; i++ {
		row := [BOARD_SIZE]*ShipDeck{}

		for j := 0; j < BOARD_SIZE; j++ {
			row[j] = nil
		}

		board[i] = row
	}

	return board
}

func RandomizeShips(board *GameBoard) {
	board[0][0] = &ShipDeck{
		isTouched: false,
		isDead:    false,
	}
}

var RowsAlphabet = [BOARD_SIZE]string{
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
	fmt.Println()

	fmt.Print("   ")

	for i := range RowsAlphabet {
		fmt.Printf("%v  ", RowsAlphabet[i])
	}

	for i, row := range board {
		//fmt.Printf("\n%v", row)
		fmt.Println()

		if i == 9 {
			fmt.Printf("%v ", i+1)
		} else {
			fmt.Printf("%v  ", i+1)
		}

		for _, value := range row {

			if value == {
				fmt.Printf("%v ", "🔥")
			} else {
				fmt.Printf("%v ", "◻ ")
				//fmt.Printf("%v  ", "■")
			}

			//🔥❌◻
		}
	}
	fmt.Println()
	fmt.Println()
}

type Cords struct {
	X uint8
	Y uint8
}

func Turn(board *GameBoard, cords Cords) error {
	x, y := cords.X, cords.Y

	if x > 10 {
		return errors.New("x coord is too big")
	}
	if y > 10 {
		return errors.New("y coord is too big")
	}

	board[y-1][x-1] = true

	return nil
}
