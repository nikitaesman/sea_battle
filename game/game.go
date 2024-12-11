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

type Field struct {
	isShip     bool
	isHit      bool
	isShipDead bool
}

type GameBoard = [BOARD_SIZE][BOARD_SIZE]*Field

func GetBoard() GameBoard {
	board := getEmptyBoard()

	PlaceAllShips(&board)

	return board
}

func getEmptyBoard() GameBoard {
	var board GameBoard = GameBoard{}

	for i := 0; i < BOARD_SIZE; i++ {
		row := [BOARD_SIZE]*Field{}

		for j := 0; j < BOARD_SIZE; j++ {
			row[j] = &Field{
				isShip:     false,
				isHit:      false,
				isShipDead: false,
			}
		}

		board[i] = row
	}

	return board
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

type Cords struct {
	X int
	Y int
}

func Turn(board *GameBoard, cords Cords) error {
	field, err := GetFieldByCords(board, cords)

	if err != nil {
		return err
	}

	if field.isHit {
		return errors.New("already existed turn")
	}

	field.isHit = true

	if !field.isShip {
		return nil
	}

	ship, err := getCompositeShip(board, cords)

	if err != nil {
		return err
	}

	if checkShipIsDead(ship) {
		makeShipDead(ship)
	}

	return nil
}

func checkShipIsDead(ship []*Field) bool {
	for _, partOfShip := range ship {
		if !partOfShip.isHit {
			return false
		}
	}

	return true
}

func makeShipDead(ship []*Field) {
	for _, partOfShip := range ship {
		partOfShip.isShipDead = true
	}
}

func getCompositeShip(board *GameBoard, cords Cords) ([]*Field, error) {
	field, err := GetFieldByCords(board, cords)

	if err != nil {
		return nil, err
	}

	if !field.isShip {
		return nil, errors.New("getCompositeShip: field should be an ship")
	}

	shipFields := []*Field{field}

	checkShifts := []Cords{
		{
			X: 0,
			Y: -1,
		},
		{
			X: 1,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: -1,
			Y: 0,
		},
	}

	for _, shift := range checkShifts {
		shipFields = getShipAppendSlice(
			board,
			cords,
			shift,
			shipFields,
		)
	}

	return shipFields, nil
}

func getShipAppendSlice(
	board *GameBoard,
	cords Cords,
	shift Cords,
	shipFields []*Field,
) []*Field {
	fmt.Println("getShipAppendSlice call")
	field, err := getFieldByShift(board, cords, shift)
	fmt.Printf("\nCords%v", cords)
	fmt.Printf("\n shift%v", shift)
	fmt.Printf("\n field link %v\n", &field)
	fmt.Printf("\n field%v\n", field)

	if err != nil {
		return shipFields
	}

	if field.isShip {
		shipFields = append(shipFields, field)

		newCords := Cords{
			X: cords.X + shift.X,
			Y: cords.Y + shift.Y,
		}

		fmt.Printf("\newCords%v", newCords)

		return getShipAppendSlice(board, newCords, shift, shipFields)
	}

	return shipFields
}

func GetFieldByCords(board *GameBoard, cords Cords) (*Field, error) {
	x, y := cords.X, cords.Y

	if x > 10 {
		return nil, errors.New("x coord is too big. Supported range 1-10")
	}
	if x < 1 {
		return nil, errors.New("x coord is too small. Supported range 1-10")
	}

	if y > 10 {
		return nil, errors.New("y coord is too big. Supported range 1-10")
	}
	if y < 1 {
		return nil, errors.New("y coord is too small. Supported range 1-10")
	}

	return board[y-1][x-1], nil
}

func getFieldByShift(board *GameBoard, cords Cords, shift Cords) (*Field, error) {
	if shift.X == 0 && shift.Y == 0 {
		return nil, errors.New("shift cords should not be zero")
	}

	cordsWithShift := Cords{
		X: cords.X + shift.X,
		Y: cords.Y + shift.Y,
	}

	return GetFieldByCords(board, cordsWithShift)
}
