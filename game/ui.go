package game

import "fmt"

func PrintGameBoard(board GameBoard, invisibleShips bool) {
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

		for _, field := range row {
			if field.isShip {
				if field.isShipDead {
					fmt.Printf("%v ", "❌")
					continue
				}

				if field.isHit {
					fmt.Printf("%v ", "🔥")
					continue
				} else if !invisibleShips {
					fmt.Printf("%v ", "& ")
					continue
				}
			} else {
				if field.isHit {
					fmt.Printf("%v  ", "■")
					continue
				}
			}

			fmt.Printf("%v ", "◻ ")

			//🔥❌◻
		}
	}
	fmt.Println()
	fmt.Println()
}
